from .base import *

ALLOWED_HOSTS = ['*']
DEBUG = True

# Database
# https://docs.djangoproject.com/en/2.0/ref/settings/#databases

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'vikkyshostak',
        'USER': 'vikkyshostak',
        'PASSWORD': '123456',
        'HOST': 'localhost',
        'PORT': '5432',
    }
}

# Email Backend
# https://docs.djangoproject.com/en/dev/topics/email/

EMAIL_BACKEND = 'django.core.mail.backends.dummy.EmailBackend'

# Django Cache framework
# https://docs.djangoproject.com/en/dev/topics/cache/

CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.dummy.DummyCache',
    }
}

# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/dev/howto/static-files/

STATICFILES_DIRS = [
    os.path.join(BASE_DIR, 'static'),
]
