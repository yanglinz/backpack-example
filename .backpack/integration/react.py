import functools
import os
import pathlib
import socket

import requests
from bs4 import BeautifulSoup

DEV_SERVER = "http://docker.for.mac.localhost:3000"
DEV_SERVER_HOST = "http://localhost:3000"


def get_dev_markup(index_url):
    try:
        resp = requests.get(index_url)
    except requests.exceptions.ConnectionError:
        # In development, there are cases e.g. running tests
        # where the frontend server may not be available.
        # In those cases, we can fallback to an empty markup
        empty_markup = "<html><head></head><body></body></html>"
        return empty_markup
    return resp.text


@functools.lru_cache()
def get_prod_markup(root_dir):
    root_path = pathlib.Path(root_dir)
    index_html = root_path.joinpath("build/index.html")
    with open(str(index_html)) as f:
        return f.read()


def get_assets(html_markup, public_host=""):
    markup_soup = BeautifulSoup(html_markup, "html.parser")

    style_tags = markup_soup.find("head").find_all("link")
    style_tags = [l for l in style_tags if l.get("rel") == ["stylesheet"]]
    for s in style_tags:
        s["href"] = [f"{public_host}{s['href']}"]

    script_tags = markup_soup.find("body").find_all("script")
    for s in script_tags:
        if s.get("src"):
            s["src"] = f"{public_host}{s['src']}"

    style_markup = "".join([str(s) for s in style_tags])
    script_markup = "".join([str(s) for s in script_tags])
    return {"cra_style_markup": style_markup, "cra_script_markup": script_markup}


def resolve_assets(
    root_dir=os.getcwd(), dev_server=False, dev_server_url=DEV_SERVER, public_host=None,
):
    """
    Resolve the path to js/css/media assets for a CRA instance
    :param root_dir: Root of the dir where package.json resides
    :param dev_server: Whether we're resolving against dev_server or pre-built production html
    :param dev_server_url: Host where index of dev server can be resolved
    :param public_host: Public host on what to prepend to asset paths
    """
    if dev_server:
        markup = get_dev_markup(dev_server_url)
    else:
        markup = get_prod_markup(root_dir)

    return get_assets(markup, public_host=public_host)


@functools.lru_cache()
def resolve_assets_cached(*args, **kwargs):
    return resolve_assets(*args, **kwargs)


@functools.lru_cache()
def is_running_dev_server():
    return os.environ.get("BACKPACK_DOCKER_COMPOSE") == "true"


def cra_context_processor(request):
    dev_server = is_running_dev_server()
    dev_server_url = DEV_SERVER
    public_host = DEV_SERVER_HOST if dev_server else ""

    if dev_server:
        return resolve_assets(
            dev_server=dev_server,
            dev_server_url=dev_server_url,
            public_host=public_host,
        )

    return resolve_assets_cached(
        dev_server=dev_server, dev_server_url=dev_server_url, public_host=public_host,
    )
