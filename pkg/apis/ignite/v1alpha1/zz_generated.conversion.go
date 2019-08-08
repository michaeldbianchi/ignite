// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	ignite "github.com/weaveworks/ignite/pkg/apis/ignite"
	metav1alpha1 "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*FileMapping)(nil), (*ignite.FileMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FileMapping_To_ignite_FileMapping(a.(*FileMapping), b.(*ignite.FileMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.FileMapping)(nil), (*FileMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_FileMapping_To_v1alpha1_FileMapping(a.(*ignite.FileMapping), b.(*FileMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Image)(nil), (*ignite.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Image_To_ignite_Image(a.(*Image), b.(*ignite.Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.Image)(nil), (*Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_Image_To_v1alpha1_Image(a.(*ignite.Image), b.(*Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ImageSpec)(nil), (*ignite.ImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec(a.(*ImageSpec), b.(*ignite.ImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.ImageSpec)(nil), (*ImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec(a.(*ignite.ImageSpec), b.(*ImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ImageStatus)(nil), (*ignite.ImageStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ImageStatus_To_ignite_ImageStatus(a.(*ImageStatus), b.(*ignite.ImageStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.ImageStatus)(nil), (*ImageStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_ImageStatus_To_v1alpha1_ImageStatus(a.(*ignite.ImageStatus), b.(*ImageStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Kernel)(nil), (*ignite.Kernel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Kernel_To_ignite_Kernel(a.(*Kernel), b.(*ignite.Kernel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.Kernel)(nil), (*Kernel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_Kernel_To_v1alpha1_Kernel(a.(*ignite.Kernel), b.(*Kernel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KernelSpec)(nil), (*ignite.KernelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec(a.(*KernelSpec), b.(*ignite.KernelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.KernelSpec)(nil), (*KernelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec(a.(*ignite.KernelSpec), b.(*KernelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KernelStatus)(nil), (*ignite.KernelStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KernelStatus_To_ignite_KernelStatus(a.(*KernelStatus), b.(*ignite.KernelStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.KernelStatus)(nil), (*KernelStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_KernelStatus_To_v1alpha1_KernelStatus(a.(*ignite.KernelStatus), b.(*KernelStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OCIImageClaim)(nil), (*ignite.OCIImageClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(a.(*OCIImageClaim), b.(*ignite.OCIImageClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.OCIImageClaim)(nil), (*OCIImageClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(a.(*ignite.OCIImageClaim), b.(*OCIImageClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OCIImageSource)(nil), (*ignite.OCIImageSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(a.(*OCIImageSource), b.(*ignite.OCIImageSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.OCIImageSource)(nil), (*OCIImageSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(a.(*ignite.OCIImageSource), b.(*OCIImageSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Pool)(nil), (*ignite.Pool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Pool_To_ignite_Pool(a.(*Pool), b.(*ignite.Pool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.Pool)(nil), (*Pool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_Pool_To_v1alpha1_Pool(a.(*ignite.Pool), b.(*Pool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PoolDevice)(nil), (*ignite.PoolDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PoolDevice_To_ignite_PoolDevice(a.(*PoolDevice), b.(*ignite.PoolDevice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.PoolDevice)(nil), (*PoolDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_PoolDevice_To_v1alpha1_PoolDevice(a.(*ignite.PoolDevice), b.(*PoolDevice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PoolSpec)(nil), (*ignite.PoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PoolSpec_To_ignite_PoolSpec(a.(*PoolSpec), b.(*ignite.PoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.PoolSpec)(nil), (*PoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_PoolSpec_To_v1alpha1_PoolSpec(a.(*ignite.PoolSpec), b.(*PoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PoolStatus)(nil), (*ignite.PoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PoolStatus_To_ignite_PoolStatus(a.(*PoolStatus), b.(*ignite.PoolStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.PoolStatus)(nil), (*PoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_PoolStatus_To_v1alpha1_PoolStatus(a.(*ignite.PoolStatus), b.(*PoolStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SSH)(nil), (*ignite.SSH)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SSH_To_ignite_SSH(a.(*SSH), b.(*ignite.SSH), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.SSH)(nil), (*SSH)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_SSH_To_v1alpha1_SSH(a.(*ignite.SSH), b.(*SSH), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VM)(nil), (*ignite.VM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VM_To_ignite_VM(a.(*VM), b.(*ignite.VM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VM)(nil), (*VM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VM_To_v1alpha1_VM(a.(*ignite.VM), b.(*VM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VMImageSpec)(nil), (*ignite.VMImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(a.(*VMImageSpec), b.(*ignite.VMImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VMImageSpec)(nil), (*VMImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(a.(*ignite.VMImageSpec), b.(*VMImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VMKernelSpec)(nil), (*ignite.VMKernelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(a.(*VMKernelSpec), b.(*ignite.VMKernelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VMKernelSpec)(nil), (*VMKernelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(a.(*ignite.VMKernelSpec), b.(*VMKernelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VMNetworkSpec)(nil), (*ignite.VMNetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(a.(*VMNetworkSpec), b.(*ignite.VMNetworkSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VMNetworkSpec)(nil), (*VMNetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec(a.(*ignite.VMNetworkSpec), b.(*VMNetworkSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VMSpec)(nil), (*ignite.VMSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMSpec_To_ignite_VMSpec(a.(*VMSpec), b.(*ignite.VMSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VMSpec)(nil), (*VMSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMSpec_To_v1alpha1_VMSpec(a.(*ignite.VMSpec), b.(*VMSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VMStatus)(nil), (*ignite.VMStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMStatus_To_ignite_VMStatus(a.(*VMStatus), b.(*ignite.VMStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ignite.VMStatus)(nil), (*VMStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMStatus_To_v1alpha1_VMStatus(a.(*ignite.VMStatus), b.(*VMStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ignite.VMSpec)(nil), (*VMSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMSpec_To_v1alpha1_VMSpec(a.(*ignite.VMSpec), b.(*VMSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ignite.VMStatus)(nil), (*VMStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ignite_VMStatus_To_v1alpha1_VMStatus(a.(*ignite.VMStatus), b.(*VMStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*VMSpec)(nil), (*ignite.VMSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMSpec_To_ignite_VMSpec(a.(*VMSpec), b.(*ignite.VMSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*VMStatus)(nil), (*ignite.VMStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VMStatus_To_ignite_VMStatus(a.(*VMStatus), b.(*ignite.VMStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_FileMapping_To_ignite_FileMapping(in *FileMapping, out *ignite.FileMapping, s conversion.Scope) error {
	out.HostPath = in.HostPath
	out.VMPath = in.VMPath
	return nil
}

// Convert_v1alpha1_FileMapping_To_ignite_FileMapping is an autogenerated conversion function.
func Convert_v1alpha1_FileMapping_To_ignite_FileMapping(in *FileMapping, out *ignite.FileMapping, s conversion.Scope) error {
	return autoConvert_v1alpha1_FileMapping_To_ignite_FileMapping(in, out, s)
}

func autoConvert_ignite_FileMapping_To_v1alpha1_FileMapping(in *ignite.FileMapping, out *FileMapping, s conversion.Scope) error {
	out.HostPath = in.HostPath
	out.VMPath = in.VMPath
	return nil
}

// Convert_ignite_FileMapping_To_v1alpha1_FileMapping is an autogenerated conversion function.
func Convert_ignite_FileMapping_To_v1alpha1_FileMapping(in *ignite.FileMapping, out *FileMapping, s conversion.Scope) error {
	return autoConvert_ignite_FileMapping_To_v1alpha1_FileMapping(in, out, s)
}

func autoConvert_v1alpha1_Image_To_ignite_Image(in *Image, out *ignite.Image, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ImageStatus_To_ignite_ImageStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Image_To_ignite_Image is an autogenerated conversion function.
func Convert_v1alpha1_Image_To_ignite_Image(in *Image, out *ignite.Image, s conversion.Scope) error {
	return autoConvert_v1alpha1_Image_To_ignite_Image(in, out, s)
}

func autoConvert_ignite_Image_To_v1alpha1_Image(in *ignite.Image, out *Image, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ignite_ImageStatus_To_v1alpha1_ImageStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_Image_To_v1alpha1_Image is an autogenerated conversion function.
func Convert_ignite_Image_To_v1alpha1_Image(in *ignite.Image, out *Image, s conversion.Scope) error {
	return autoConvert_ignite_Image_To_v1alpha1_Image(in, out, s)
}

func autoConvert_v1alpha1_ImageSpec_To_ignite_ImageSpec(in *ImageSpec, out *ignite.ImageSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec is an autogenerated conversion function.
func Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec(in *ImageSpec, out *ignite.ImageSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ImageSpec_To_ignite_ImageSpec(in, out, s)
}

func autoConvert_ignite_ImageSpec_To_v1alpha1_ImageSpec(in *ignite.ImageSpec, out *ImageSpec, s conversion.Scope) error {
	if err := Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec is an autogenerated conversion function.
func Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec(in *ignite.ImageSpec, out *ImageSpec, s conversion.Scope) error {
	return autoConvert_ignite_ImageSpec_To_v1alpha1_ImageSpec(in, out, s)
}

func autoConvert_v1alpha1_ImageStatus_To_ignite_ImageStatus(in *ImageStatus, out *ignite.ImageStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(&in.OCISource, &out.OCISource, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ImageStatus_To_ignite_ImageStatus is an autogenerated conversion function.
func Convert_v1alpha1_ImageStatus_To_ignite_ImageStatus(in *ImageStatus, out *ignite.ImageStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ImageStatus_To_ignite_ImageStatus(in, out, s)
}

func autoConvert_ignite_ImageStatus_To_v1alpha1_ImageStatus(in *ignite.ImageStatus, out *ImageStatus, s conversion.Scope) error {
	if err := Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(&in.OCISource, &out.OCISource, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_ImageStatus_To_v1alpha1_ImageStatus is an autogenerated conversion function.
func Convert_ignite_ImageStatus_To_v1alpha1_ImageStatus(in *ignite.ImageStatus, out *ImageStatus, s conversion.Scope) error {
	return autoConvert_ignite_ImageStatus_To_v1alpha1_ImageStatus(in, out, s)
}

func autoConvert_v1alpha1_Kernel_To_ignite_Kernel(in *Kernel, out *ignite.Kernel, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KernelStatus_To_ignite_KernelStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Kernel_To_ignite_Kernel is an autogenerated conversion function.
func Convert_v1alpha1_Kernel_To_ignite_Kernel(in *Kernel, out *ignite.Kernel, s conversion.Scope) error {
	return autoConvert_v1alpha1_Kernel_To_ignite_Kernel(in, out, s)
}

func autoConvert_ignite_Kernel_To_v1alpha1_Kernel(in *ignite.Kernel, out *Kernel, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ignite_KernelStatus_To_v1alpha1_KernelStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_Kernel_To_v1alpha1_Kernel is an autogenerated conversion function.
func Convert_ignite_Kernel_To_v1alpha1_Kernel(in *ignite.Kernel, out *Kernel, s conversion.Scope) error {
	return autoConvert_ignite_Kernel_To_v1alpha1_Kernel(in, out, s)
}

func autoConvert_v1alpha1_KernelSpec_To_ignite_KernelSpec(in *KernelSpec, out *ignite.KernelSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec is an autogenerated conversion function.
func Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec(in *KernelSpec, out *ignite.KernelSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_KernelSpec_To_ignite_KernelSpec(in, out, s)
}

func autoConvert_ignite_KernelSpec_To_v1alpha1_KernelSpec(in *ignite.KernelSpec, out *KernelSpec, s conversion.Scope) error {
	if err := Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec is an autogenerated conversion function.
func Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec(in *ignite.KernelSpec, out *KernelSpec, s conversion.Scope) error {
	return autoConvert_ignite_KernelSpec_To_v1alpha1_KernelSpec(in, out, s)
}

func autoConvert_v1alpha1_KernelStatus_To_ignite_KernelStatus(in *KernelStatus, out *ignite.KernelStatus, s conversion.Scope) error {
	out.Version = in.Version
	if err := Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(&in.OCISource, &out.OCISource, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_KernelStatus_To_ignite_KernelStatus is an autogenerated conversion function.
func Convert_v1alpha1_KernelStatus_To_ignite_KernelStatus(in *KernelStatus, out *ignite.KernelStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_KernelStatus_To_ignite_KernelStatus(in, out, s)
}

func autoConvert_ignite_KernelStatus_To_v1alpha1_KernelStatus(in *ignite.KernelStatus, out *KernelStatus, s conversion.Scope) error {
	out.Version = in.Version
	if err := Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(&in.OCISource, &out.OCISource, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_KernelStatus_To_v1alpha1_KernelStatus is an autogenerated conversion function.
func Convert_ignite_KernelStatus_To_v1alpha1_KernelStatus(in *ignite.KernelStatus, out *KernelStatus, s conversion.Scope) error {
	return autoConvert_ignite_KernelStatus_To_v1alpha1_KernelStatus(in, out, s)
}

func autoConvert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(in *OCIImageClaim, out *ignite.OCIImageClaim, s conversion.Scope) error {
	out.Type = ignite.ImageSourceType(in.Type)
	out.Ref = metav1alpha1.OCIImageRef(in.Ref)
	return nil
}

// Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim is an autogenerated conversion function.
func Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(in *OCIImageClaim, out *ignite.OCIImageClaim, s conversion.Scope) error {
	return autoConvert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(in, out, s)
}

func autoConvert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(in *ignite.OCIImageClaim, out *OCIImageClaim, s conversion.Scope) error {
	out.Type = ImageSourceType(in.Type)
	out.Ref = metav1alpha1.OCIImageRef(in.Ref)
	return nil
}

// Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim is an autogenerated conversion function.
func Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(in *ignite.OCIImageClaim, out *OCIImageClaim, s conversion.Scope) error {
	return autoConvert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(in, out, s)
}

func autoConvert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(in *OCIImageSource, out *ignite.OCIImageSource, s conversion.Scope) error {
	out.ID = in.ID
	out.Size = in.Size
	out.RepoDigests = *(*[]string)(unsafe.Pointer(&in.RepoDigests))
	return nil
}

// Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource is an autogenerated conversion function.
func Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(in *OCIImageSource, out *ignite.OCIImageSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(in, out, s)
}

func autoConvert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(in *ignite.OCIImageSource, out *OCIImageSource, s conversion.Scope) error {
	out.ID = in.ID
	out.Size = in.Size
	out.RepoDigests = *(*[]string)(unsafe.Pointer(&in.RepoDigests))
	return nil
}

// Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource is an autogenerated conversion function.
func Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(in *ignite.OCIImageSource, out *OCIImageSource, s conversion.Scope) error {
	return autoConvert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(in, out, s)
}

func autoConvert_v1alpha1_Pool_To_ignite_Pool(in *Pool, out *ignite.Pool, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	if err := Convert_v1alpha1_PoolSpec_To_ignite_PoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PoolStatus_To_ignite_PoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Pool_To_ignite_Pool is an autogenerated conversion function.
func Convert_v1alpha1_Pool_To_ignite_Pool(in *Pool, out *ignite.Pool, s conversion.Scope) error {
	return autoConvert_v1alpha1_Pool_To_ignite_Pool(in, out, s)
}

func autoConvert_ignite_Pool_To_v1alpha1_Pool(in *ignite.Pool, out *Pool, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	if err := Convert_ignite_PoolSpec_To_v1alpha1_PoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ignite_PoolStatus_To_v1alpha1_PoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_Pool_To_v1alpha1_Pool is an autogenerated conversion function.
func Convert_ignite_Pool_To_v1alpha1_Pool(in *ignite.Pool, out *Pool, s conversion.Scope) error {
	return autoConvert_ignite_Pool_To_v1alpha1_Pool(in, out, s)
}

func autoConvert_v1alpha1_PoolDevice_To_ignite_PoolDevice(in *PoolDevice, out *ignite.PoolDevice, s conversion.Scope) error {
	out.Size = in.Size
	out.Parent = in.Parent
	out.Type = ignite.PoolDeviceType(in.Type)
	out.MetadataPath = in.MetadataPath
	return nil
}

// Convert_v1alpha1_PoolDevice_To_ignite_PoolDevice is an autogenerated conversion function.
func Convert_v1alpha1_PoolDevice_To_ignite_PoolDevice(in *PoolDevice, out *ignite.PoolDevice, s conversion.Scope) error {
	return autoConvert_v1alpha1_PoolDevice_To_ignite_PoolDevice(in, out, s)
}

func autoConvert_ignite_PoolDevice_To_v1alpha1_PoolDevice(in *ignite.PoolDevice, out *PoolDevice, s conversion.Scope) error {
	out.Size = in.Size
	out.Parent = in.Parent
	out.Type = PoolDeviceType(in.Type)
	out.MetadataPath = in.MetadataPath
	return nil
}

// Convert_ignite_PoolDevice_To_v1alpha1_PoolDevice is an autogenerated conversion function.
func Convert_ignite_PoolDevice_To_v1alpha1_PoolDevice(in *ignite.PoolDevice, out *PoolDevice, s conversion.Scope) error {
	return autoConvert_ignite_PoolDevice_To_v1alpha1_PoolDevice(in, out, s)
}

func autoConvert_v1alpha1_PoolSpec_To_ignite_PoolSpec(in *PoolSpec, out *ignite.PoolSpec, s conversion.Scope) error {
	out.MetadataSize = in.MetadataSize
	out.DataSize = in.DataSize
	out.AllocationSize = in.AllocationSize
	out.MetadataPath = in.MetadataPath
	out.DataPath = in.DataPath
	return nil
}

// Convert_v1alpha1_PoolSpec_To_ignite_PoolSpec is an autogenerated conversion function.
func Convert_v1alpha1_PoolSpec_To_ignite_PoolSpec(in *PoolSpec, out *ignite.PoolSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PoolSpec_To_ignite_PoolSpec(in, out, s)
}

func autoConvert_ignite_PoolSpec_To_v1alpha1_PoolSpec(in *ignite.PoolSpec, out *PoolSpec, s conversion.Scope) error {
	out.MetadataSize = in.MetadataSize
	out.DataSize = in.DataSize
	out.AllocationSize = in.AllocationSize
	out.MetadataPath = in.MetadataPath
	out.DataPath = in.DataPath
	return nil
}

// Convert_ignite_PoolSpec_To_v1alpha1_PoolSpec is an autogenerated conversion function.
func Convert_ignite_PoolSpec_To_v1alpha1_PoolSpec(in *ignite.PoolSpec, out *PoolSpec, s conversion.Scope) error {
	return autoConvert_ignite_PoolSpec_To_v1alpha1_PoolSpec(in, out, s)
}

func autoConvert_v1alpha1_PoolStatus_To_ignite_PoolStatus(in *PoolStatus, out *ignite.PoolStatus, s conversion.Scope) error {
	out.Devices = *(*[]*ignite.PoolDevice)(unsafe.Pointer(&in.Devices))
	return nil
}

// Convert_v1alpha1_PoolStatus_To_ignite_PoolStatus is an autogenerated conversion function.
func Convert_v1alpha1_PoolStatus_To_ignite_PoolStatus(in *PoolStatus, out *ignite.PoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PoolStatus_To_ignite_PoolStatus(in, out, s)
}

func autoConvert_ignite_PoolStatus_To_v1alpha1_PoolStatus(in *ignite.PoolStatus, out *PoolStatus, s conversion.Scope) error {
	out.Devices = *(*[]*PoolDevice)(unsafe.Pointer(&in.Devices))
	return nil
}

// Convert_ignite_PoolStatus_To_v1alpha1_PoolStatus is an autogenerated conversion function.
func Convert_ignite_PoolStatus_To_v1alpha1_PoolStatus(in *ignite.PoolStatus, out *PoolStatus, s conversion.Scope) error {
	return autoConvert_ignite_PoolStatus_To_v1alpha1_PoolStatus(in, out, s)
}

func autoConvert_v1alpha1_SSH_To_ignite_SSH(in *SSH, out *ignite.SSH, s conversion.Scope) error {
	out.Generate = in.Generate
	out.PublicKey = in.PublicKey
	return nil
}

// Convert_v1alpha1_SSH_To_ignite_SSH is an autogenerated conversion function.
func Convert_v1alpha1_SSH_To_ignite_SSH(in *SSH, out *ignite.SSH, s conversion.Scope) error {
	return autoConvert_v1alpha1_SSH_To_ignite_SSH(in, out, s)
}

func autoConvert_ignite_SSH_To_v1alpha1_SSH(in *ignite.SSH, out *SSH, s conversion.Scope) error {
	out.Generate = in.Generate
	out.PublicKey = in.PublicKey
	return nil
}

// Convert_ignite_SSH_To_v1alpha1_SSH is an autogenerated conversion function.
func Convert_ignite_SSH_To_v1alpha1_SSH(in *ignite.SSH, out *SSH, s conversion.Scope) error {
	return autoConvert_ignite_SSH_To_v1alpha1_SSH(in, out, s)
}

func autoConvert_v1alpha1_VM_To_ignite_VM(in *VM, out *ignite.VM, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VMSpec_To_ignite_VMSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VMStatus_To_ignite_VMStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VM_To_ignite_VM is an autogenerated conversion function.
func Convert_v1alpha1_VM_To_ignite_VM(in *VM, out *ignite.VM, s conversion.Scope) error {
	return autoConvert_v1alpha1_VM_To_ignite_VM(in, out, s)
}

func autoConvert_ignite_VM_To_v1alpha1_VM(in *ignite.VM, out *VM, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ignite_VMSpec_To_v1alpha1_VMSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ignite_VMStatus_To_v1alpha1_VMStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_VM_To_v1alpha1_VM is an autogenerated conversion function.
func Convert_ignite_VM_To_v1alpha1_VM(in *ignite.VM, out *VM, s conversion.Scope) error {
	return autoConvert_ignite_VM_To_v1alpha1_VM(in, out, s)
}

func autoConvert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(in *VMImageSpec, out *ignite.VMImageSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec is an autogenerated conversion function.
func Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(in *VMImageSpec, out *ignite.VMImageSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(in, out, s)
}

func autoConvert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(in *ignite.VMImageSpec, out *VMImageSpec, s conversion.Scope) error {
	if err := Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	return nil
}

// Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec is an autogenerated conversion function.
func Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(in *ignite.VMImageSpec, out *VMImageSpec, s conversion.Scope) error {
	return autoConvert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(in, out, s)
}

func autoConvert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(in *VMKernelSpec, out *ignite.VMKernelSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_OCIImageClaim_To_ignite_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	out.CmdLine = in.CmdLine
	return nil
}

// Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec is an autogenerated conversion function.
func Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(in *VMKernelSpec, out *ignite.VMKernelSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(in, out, s)
}

func autoConvert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(in *ignite.VMKernelSpec, out *VMKernelSpec, s conversion.Scope) error {
	if err := Convert_ignite_OCIImageClaim_To_v1alpha1_OCIImageClaim(&in.OCIClaim, &out.OCIClaim, s); err != nil {
		return err
	}
	out.CmdLine = in.CmdLine
	return nil
}

// Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec is an autogenerated conversion function.
func Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(in *ignite.VMKernelSpec, out *VMKernelSpec, s conversion.Scope) error {
	return autoConvert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(in, out, s)
}

func autoConvert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(in *VMNetworkSpec, out *ignite.VMNetworkSpec, s conversion.Scope) error {
	out.Mode = ignite.NetworkMode(in.Mode)
	out.Ports = *(*metav1alpha1.PortMappings)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec is an autogenerated conversion function.
func Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(in *VMNetworkSpec, out *ignite.VMNetworkSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(in, out, s)
}

func autoConvert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec(in *ignite.VMNetworkSpec, out *VMNetworkSpec, s conversion.Scope) error {
	out.Mode = NetworkMode(in.Mode)
	out.Ports = *(*metav1alpha1.PortMappings)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec is an autogenerated conversion function.
func Convert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec(in *ignite.VMNetworkSpec, out *VMNetworkSpec, s conversion.Scope) error {
	return autoConvert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec(in, out, s)
}

func autoConvert_v1alpha1_VMSpec_To_ignite_VMSpec(in *VMSpec, out *ignite.VMSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(&in.Kernel, &out.Kernel, s); err != nil {
		return err
	}
	out.CPUs = in.CPUs
	out.Memory = in.Memory
	out.DiskSize = in.DiskSize
	if err := Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(&in.Network, &out.Network, s); err != nil {
		return err
	}
	out.CopyFiles = *(*[]ignite.FileMapping)(unsafe.Pointer(&in.CopyFiles))
	out.SSH = (*ignite.SSH)(unsafe.Pointer(in.SSH))
	return nil
}

func autoConvert_ignite_VMSpec_To_v1alpha1_VMSpec(in *ignite.VMSpec, out *VMSpec, s conversion.Scope) error {
	if err := Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(&in.Kernel, &out.Kernel, s); err != nil {
		return err
	}
	out.CPUs = in.CPUs
	out.Memory = in.Memory
	out.DiskSize = in.DiskSize
	if err := Convert_ignite_VMNetworkSpec_To_v1alpha1_VMNetworkSpec(&in.Network, &out.Network, s); err != nil {
		return err
	}
	// WARNING: in.Storage requires manual conversion: does not exist in peer-type
	out.CopyFiles = *(*[]FileMapping)(unsafe.Pointer(&in.CopyFiles))
	out.SSH = (*SSH)(unsafe.Pointer(in.SSH))
	return nil
}

func autoConvert_v1alpha1_VMStatus_To_ignite_VMStatus(in *VMStatus, out *ignite.VMStatus, s conversion.Scope) error {
	// WARNING: in.State requires manual conversion: does not exist in peer-type
	out.IPAddresses = *(*metav1alpha1.IPAddresses)(unsafe.Pointer(&in.IPAddresses))
	if err := Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(&in.Kernel, &out.Kernel, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_ignite_VMStatus_To_v1alpha1_VMStatus(in *ignite.VMStatus, out *VMStatus, s conversion.Scope) error {
	// WARNING: in.Running requires manual conversion: does not exist in peer-type
	// WARNING: in.Runtime requires manual conversion: does not exist in peer-type
	out.IPAddresses = *(*metav1alpha1.IPAddresses)(unsafe.Pointer(&in.IPAddresses))
	if err := Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(&in.Kernel, &out.Kernel, s); err != nil {
		return err
	}
	return nil
}
