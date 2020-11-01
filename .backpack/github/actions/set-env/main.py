import os
from os import path

from yaml import load


def export_env(key, value):
    with open(os.environ["GITHUB_ENV"], "a") as f:
        print(f"exporting {key}={value}")
        f.write(f"{key}={value}\n")


def get_backpack_manifest():
    manifest_path = path.join(os.getcwd(), "backpack.yml")
    with open(manifest_path, "r") as f:
        content = f.read()
        return load(content)


def export_backpack_envs(manifest):
    # TODO: Namespace vars
    export_env("APP_NAME", manifest["name"])
    export_env("GCP_PROJECT_ID", "default-263000")
    export_env("RUNTIME_PLATFORM", manifest["runtime"])
    project_name = manifest["projects"][0]["name"]
    export_env("BACKPACK_DOCKER_SERVICE_NAME", f"{project_name}_server")


if __name__ == "__main__":
    manifest = get_backpack_manifest()
    export_backpack_envs(manifest)
