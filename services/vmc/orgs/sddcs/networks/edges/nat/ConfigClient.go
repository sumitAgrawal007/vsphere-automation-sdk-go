/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Config
 * Used by client-side stubs.
 */

package nat

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ConfigClient interface {

    // Delete all NAT configuration for the specified management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Delete(orgParam string, sddcParam string, edgeIdParam string) error

    // Retrieve NAT configuration for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @return com.vmware.vmc.model.Nat
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeIdParam string) (model.Nat, error)

    // Modify NAT configuration for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param natParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Update(orgParam string, sddcParam string, edgeIdParam string, natParam model.Nat) error
}
