from django.db import models

class Blog(models.Model):
    body = models.CharField(max_length=50)
