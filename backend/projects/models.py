from django.db import models
from .choises import STATUSES


class Project(models.Model):
    """
    Class Project
    """

    name = models.CharField(max_length=255)
    status = models.CharField(max_length=255, choices=STATUSES, default=0)

    class Meta:
        ordering = ['id']
        verbose_name = 'Project'
        verbose_name_plural = 'Projects'

    def __str__(self):
        return f'{self.name} ({self.pk})'
