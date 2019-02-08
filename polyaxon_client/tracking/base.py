# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import atexit
import sys
import time

from polystores.stores.manager import StoreManager

from polyaxon_client import PolyaxonClient, settings
from polyaxon_client.exceptions import PolyaxonClientException
from polyaxon_client.tracking.paths import get_outputs_path
from polyaxon_client.tracking.utils.project import get_project_info


class BaseTracker(object):
    def __init__(self,
                 project=None,
                 client=None,
                 track_logs=True,
                 track_code=True,
                 track_env=True,
                 outputs_store=None):
        if settings.NO_OP:
            return

        if not settings.IN_CLUSTER and project is None:
            raise PolyaxonClientException('Please provide a valid project.')

        self.last_status = None
        self.client = client or PolyaxonClient()
        if settings.IN_CLUSTER:
            self.user = None
        else:
            self.user = (self.client.auth.get_user().username
                         if self.client.api_config.schema_response
                         else self.client.auth.get_user().get('username'))

        username, project_name = get_project_info(current_user=self.user, project=project)
        self.track_logs = track_logs
        self.track_code = track_code
        self.track_env = track_env
        self.project = project
        self.username = username
        self.project_name = project_name
        self.outputs_store = outputs_store

        # Setup the outputs store
        if outputs_store is None and settings.IN_CLUSTER:
            self.set_outputs_store(outputs_path=get_outputs_path(), set_env_vars=True)

    def _set_health_url(self):
        raise NotImplementedError

    def log_status(self, status, message=None, traceback=None):
        raise NotImplementedError

    def _start(self):
        if settings.NO_OP:
            return

        atexit.register(self._end)
        self.start()

        def excepthook(exception, value, tb):
            self.failed(message='Type: {}, Value: {}'.format(exception, value))
            # Resume normal work
            sys.__excepthook__(exception, value, tb)

        sys.excepthook = excepthook

    def _end(self):
        if settings.NO_OP:
            return

        self.succeeded()

    def start(self):
        if settings.NO_OP:
            return

        self.log_status('running')
        self.last_status = 'running'

    def end(self, status, message=None):
        if settings.NO_OP:
            return

        if self.last_status in ['succeeded', 'failed', 'stopped']:
            return
        self.log_status(status, message)
        self.last_status = status
        time.sleep(0.1)  # Just to give the opportunity to the worker to pick the message

    def succeeded(self):
        if settings.NO_OP:
            return

        self.end('succeeded')

    def stop(self):
        if settings.NO_OP:
            return

        self.end('stopped')

    def failed(self, message=None):
        if settings.NO_OP:
            return

        self.end(status='failed', message=message)

    def set_outputs_store(self, outputs_store=None, outputs_path=None, set_env_vars=False):
        if settings.NO_OP:
            return

        if not any([outputs_store, outputs_path]):
            raise PolyaxonClientException(
                'An Store instance or and outputs path is required.')
        self.outputs_store = outputs_store or StoreManager(path=outputs_path)
        if self.outputs_store and set_env_vars:
            self.outputs_store.set_env_vars()

    def log_output(self, filename, **kwargs):
        if settings.NO_OP:
            return

        self.outputs_store.upload_file(filename=filename)

    def log_outputs(self, dirname, **kwargs):
        if settings.NO_OP:
            return

        self.outputs_store.upload_dir(dirname=dirname)
