package helpers

import (
	corev1 "k8s.io/api/core/v1"
)

func GetVolumeSourceType(source corev1.VolumeSource) string {
	if source.HostPath != nil {
		return "hostPath"
	} else if source.EmptyDir != nil {
		return "emptyDir"
	} else if source.GCEPersistentDisk != nil {
		return "gcePersistentDisk"
	} else if source.AWSElasticBlockStore != nil {
		return "awsElasticBlockStore"
	} else if source.GitRepo != nil {
		return "gitRepo"
	} else if source.Secret != nil {
		return "secret"
	} else if source.NFS != nil {
		return "nfs"
	} else if source.ISCSI != nil {
		return "iscsi"
	} else if source.Glusterfs != nil {
		return "glusterfs"
	} else if source.PersistentVolumeClaim != nil {
		return "persistentVolumeClaim"
	} else if source.RBD != nil {
		return "rbd"
	} else if source.FlexVolume != nil {
		return "flexVolume"
	} else if source.Cinder != nil {
		return "cinder"
	} else if source.CephFS != nil {
		return "cephfs"
	} else if source.Flocker != nil {
		return "flocker"
	} else if source.DownwardAPI != nil {
		return "downwardAPI"
	} else if source.FC != nil {
		return "fc"
	} else if source.AzureFile != nil {
		return "azureFile"
	} else if source.ConfigMap != nil {
		return "configMap"
	} else if source.VsphereVolume != nil {
		return "vsphereVolume"
	} else if source.Quobyte != nil {
		return "quobyte"
	} else if source.AzureDisk != nil {
		return "azureDisk"
	} else if source.PhotonPersistentDisk != nil {
		return "photonPersistentDisk"
	} else if source.Projected != nil {
		return "projected"
	} else if source.PortworxVolume != nil {
		return "portworxVolume"
	} else if source.ScaleIO != nil {
		return "scaleIO"
	} else if source.StorageOS != nil {
		return "storageOS"
	} else if source.CSI != nil {
		return "csi"
	}
	return ""
}

func GetPersistentVolumeSourceType(source corev1.PersistentVolumeSource) string {
	if source.HostPath != nil {
		return "hostPath"
	} else if source.GCEPersistentDisk != nil {
		return "gcePersistentDisk"
	} else if source.AWSElasticBlockStore != nil {
		return "awsElasticBlockStore"
	} else if source.NFS != nil {
		return "nfs"
	} else if source.ISCSI != nil {
		return "iscsi"
	} else if source.Glusterfs != nil {
		return "glusterfs"
	} else if source.RBD != nil {
		return "rbd"
	} else if source.FlexVolume != nil {
		return "flexVolume"
	} else if source.Cinder != nil {
		return "cinder"
	} else if source.CephFS != nil {
		return "cephfs"
	} else if source.Flocker != nil {
		return "flocker"
	} else if source.FC != nil {
		return "fc"
	} else if source.AzureFile != nil {
		return "azureFile"
	} else if source.VsphereVolume != nil {
		return "vsphereVolume"
	} else if source.Quobyte != nil {
		return "quobyte"
	} else if source.AzureDisk != nil {
		return "azureDisk"
	} else if source.PhotonPersistentDisk != nil {
		return "photonPersistentDisk"
	} else if source.PortworxVolume != nil {
		return "portworxVolume"
	} else if source.ScaleIO != nil {
		return "scaleIO"
	} else if source.Local != nil {
		return "local"
	} else if source.StorageOS != nil {
		return "storageOS"
	} else if source.CSI != nil {
		return "csi"
	}
	return ""
}
