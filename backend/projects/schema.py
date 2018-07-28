import graphene
from graphene_django.types import DjangoObjectType
from .models import Project


class ProjectType(DjangoObjectType):
    class Meta:
        model = Project
        interfaces = (graphene.Node,)


class Query(graphene.ObjectType):
    all_projects = graphene.List(ProjectType)

    def resolve_all_projects(self, info, **kwargs):
        return Project.objects.all()
