from django.contrib import admin
from django.urls import path
from django.views.decorators.csrf import csrf_exempt
from graphene_django.views import GraphQLView

urlpatterns = [
    path('mgmt/', admin.site.urls),
    path('__graphiql__/', GraphQLView.as_view(graphiql=True)),
    path('__gql__/', csrf_exempt(GraphQLView.as_view(batch=True))),
]
