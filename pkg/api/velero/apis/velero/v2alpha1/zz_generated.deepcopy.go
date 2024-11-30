//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSISnapshotSpec) DeepCopyInto(out *CSISnapshotSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSISnapshotSpec.
func (in *CSISnapshotSpec) DeepCopy() *CSISnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(CSISnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDownload) DeepCopyInto(out *DataDownload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDownload.
func (in *DataDownload) DeepCopy() *DataDownload {
	if in == nil {
		return nil
	}
	out := new(DataDownload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataDownload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDownloadList) DeepCopyInto(out *DataDownloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataDownload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDownloadList.
func (in *DataDownloadList) DeepCopy() *DataDownloadList {
	if in == nil {
		return nil
	}
	out := new(DataDownloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataDownloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDownloadSpec) DeepCopyInto(out *DataDownloadSpec) {
	*out = *in
	out.TargetVolume = in.TargetVolume
	if in.DataMoverConfig != nil {
		in, out := &in.DataMoverConfig, &out.DataMoverConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.OperationTimeout = in.OperationTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDownloadSpec.
func (in *DataDownloadSpec) DeepCopy() *DataDownloadSpec {
	if in == nil {
		return nil
	}
	out := new(DataDownloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDownloadStatus) DeepCopyInto(out *DataDownloadStatus) {
	*out = *in
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
	out.Progress = in.Progress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDownloadStatus.
func (in *DataDownloadStatus) DeepCopy() *DataDownloadStatus {
	if in == nil {
		return nil
	}
	out := new(DataDownloadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataUpload) DeepCopyInto(out *DataUpload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataUpload.
func (in *DataUpload) DeepCopy() *DataUpload {
	if in == nil {
		return nil
	}
	out := new(DataUpload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataUpload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataUploadList) DeepCopyInto(out *DataUploadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataUpload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataUploadList.
func (in *DataUploadList) DeepCopy() *DataUploadList {
	if in == nil {
		return nil
	}
	out := new(DataUploadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataUploadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataUploadResult) DeepCopyInto(out *DataUploadResult) {
	*out = *in
	if in.DataMoverResult != nil {
		in, out := &in.DataMoverResult, &out.DataMoverResult
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataUploadResult.
func (in *DataUploadResult) DeepCopy() *DataUploadResult {
	if in == nil {
		return nil
	}
	out := new(DataUploadResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataUploadSpec) DeepCopyInto(out *DataUploadSpec) {
	*out = *in
	if in.CSISnapshot != nil {
		in, out := &in.CSISnapshot, &out.CSISnapshot
		*out = new(CSISnapshotSpec)
		**out = **in
	}
	if in.DataMoverConfig != nil {
		in, out := &in.DataMoverConfig, &out.DataMoverConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.OperationTimeout = in.OperationTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataUploadSpec.
func (in *DataUploadSpec) DeepCopy() *DataUploadSpec {
	if in == nil {
		return nil
	}
	out := new(DataUploadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataUploadStatus) DeepCopyInto(out *DataUploadStatus) {
	*out = *in
	if in.DataMoverResult != nil {
		in, out := &in.DataMoverResult, &out.DataMoverResult
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
	out.Progress = in.Progress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataUploadStatus.
func (in *DataUploadStatus) DeepCopy() *DataUploadStatus {
	if in == nil {
		return nil
	}
	out := new(DataUploadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetVolumeSpec) DeepCopyInto(out *TargetVolumeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetVolumeSpec.
func (in *TargetVolumeSpec) DeepCopy() *TargetVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(TargetVolumeSpec)
	in.DeepCopyInto(out)
	return out
}