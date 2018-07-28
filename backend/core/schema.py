import graphene
from projects.schema import Query as ProjectQuery


class Query(ProjectQuery, graphene.ObjectType):
    pass


schema = graphene.Schema(query=Query)
