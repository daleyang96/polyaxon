#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from typing import Dict

from polyaxon.k8s.custom_resources.crd import get_custom_object

KIND = "Operation"
PLURAL = "operations"
API_VERSION = "v1"
GROUP = "core.polyaxon.com"


def get_operation_custom_object(
    resource_name: str,
    namespace: str,
    custom_object: Dict,
    annotations: Dict[str, str],
    labels: Dict[str, str],
) -> Dict:
    return get_custom_object(
        resource_name=resource_name,
        namespace=namespace,
        kind=KIND,
        api_version="{}/{}".format(GROUP, API_VERSION),
        labels=labels,
        annotations=annotations,
        custom_object=custom_object,
    )


def get_run_instance(owner: str, project: str, run_uuid: str) -> str:
    return "{}.{}.runs.{}".format(owner, project, run_uuid)


def get_notifier_instance(owner: str, project: str, run_uuid: str) -> str:
    return "{}.{}.notifiers.{}".format(owner, project, run_uuid)


def get_watchdog_instance(owner: str, project: str, run_uuid: str) -> str:
    return "{}.{}.watchdog.{}".format(owner, project, run_uuid)


def get_tuner_instance(owner: str, project: str, run_uuid: str) -> str:
    return "{}.{}.tuners.{}".format(owner, project, run_uuid)


def get_cleaner_instance(owner: str, project: str, run_uuid: str) -> str:
    return "{}.{}.cleaners.{}".format(owner, project, run_uuid)


def get_resource_name(run_uuid: str) -> str:
    return "plx-operation-{}".format(run_uuid)


def get_notifier_resource_name(run_uuid: str) -> str:
    return "plx-notifier-{}".format(run_uuid)


def get_watchdog_resource_name(run_uuid: str) -> str:
    return "plx-watchdog-{}".format(run_uuid)


def get_tuner_resource_name(run_uuid: str) -> str:
    return "plx-tuner-{}".format(run_uuid)


def get_cleaner_resource_name(run_uuid: str) -> str:
    return "plx-cleaner-{}".format(run_uuid)


def get_resource_name_for_kind(run_uuid: str, run_kind: str = None) -> str:
    if run_kind == "cleaner":
        return get_cleaner_resource_name(run_uuid)
    if run_kind == "tuner":
        return get_tuner_resource_name(run_uuid)
    if run_kind == "watchdog":
        return get_watchdog_resource_name(run_uuid)
    if run_kind == "notifier":
        return get_notifier_resource_name(run_uuid)
    # Operation
    return get_resource_name(run_uuid)
