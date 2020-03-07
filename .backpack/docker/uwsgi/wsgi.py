import os

from django.core.wsgi import get_wsgi_application

assert os.environ.get("DJANGO_SETTINGS_MODULE")
application = get_wsgi_application()
