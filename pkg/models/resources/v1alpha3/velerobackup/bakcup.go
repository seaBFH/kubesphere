/*
Copyright 2019 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package velerobackup

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/dynamicinformer"

	velerov1 "kubesphere.io/kubesphere/pkg/api/velero/apis/velero/v1"

	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	"kubesphere.io/kubesphere/pkg/models/resources/v1alpha3"
)

type veleroBackupGetter struct {
	dif dynamicinformer.DynamicSharedInformerFactory
}

func New(sharedInformers dynamicinformer.DynamicSharedInformerFactory) v1alpha3.Interface {
	return &veleroBackupGetter{dif: sharedInformers}
}

func (d *veleroBackupGetter) Get(_, name string) (runtime.Object, error) {
	// use dynamic client to get workspace
	return d.dif.ForResource(velerov1.SchemeGroupVersion.WithResource("backups")).Lister().Get(name)
	//return d.sharedInformers.Dynamic.V1alpha1().Workspaces().Lister().Get(name)
}

func convertVeleroUnstructuredToBackupList(unstructuredList []*unstructured.Unstructured) ([]*velerov1.Backup, error) {
	var backups []*velerov1.Backup
	for _, unstructured := range unstructuredList {
		backup, err := convertVeleroUnstructuredToBackup(unstructured)
		if err != nil {
			return nil, err
		}
		backups = append(backups, backup)
	}
	return backups, nil
}

func (d *veleroBackupGetter) List(_ string, query *query.Query) (*api.ListResult, error) {

	backups, err := d.dif.ForResource(velerov1.SchemeGroupVersion.WithResource("backups")).Lister().List(query.Selector())
	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, backup := range backups {
		result = append(result, backup)
	}
	listRes := v1alpha3.DefaultList(result, query, d.compare, d.filter)
	return listRes, err
}

func convertVeleroUnstructuredToBackup(unstructured *unstructured.Unstructured) (*velerov1.Backup, error) {
	backup := &velerov1.Backup{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured.Object, backup)
	if err != nil {
		return nil, err
	}
	return backup, nil
}

func (d *veleroBackupGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {

	leftBackupUnst, ok := left.(*unstructured.Unstructured)
	if !ok {
		return false
	}
	leftBackup, err := convertVeleroUnstructuredToBackup(leftBackupUnst)
	if err != nil {
		return false
	}

	rightBackupUnst, ok := right.(*unstructured.Unstructured)
	if !ok {
		return false
	}
	rightBackup, err := convertVeleroUnstructuredToBackup(rightBackupUnst)
	if err != nil {
		return false
	}

	return v1alpha3.DefaultObjectMetaCompare(leftBackup.ObjectMeta, rightBackup.ObjectMeta, field)
}

func (d *veleroBackupGetter) filter(object runtime.Object, filter query.Filter) bool {
	// convert to unstructured first
	backupUnst, ok := object.(*unstructured.Unstructured)
	if !ok {
		return false
	}
	// convert to backup
	backup, err := convertVeleroUnstructuredToBackup(backupUnst)
	if err != nil {
		return false
	}

	return v1alpha3.DefaultObjectMetaFilter(backup.ObjectMeta, filter)
}
