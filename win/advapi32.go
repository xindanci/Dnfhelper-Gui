package win

import (
	"golang.org/x/sys/windows"
)

func OpenSCManagerW(lpMachineName *uint16, lpDatabaseName *uint16, dwDesiredAccess uint32) (uintptr, error) {
	manager, err := windows.OpenSCManager(lpMachineName, lpDatabaseName, dwDesiredAccess)
	if err != nil && err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return uintptr(manager), nil
}

func CreateServiceW(scManager uintptr, serviceName *uint16, databaseName *uint16, desiredAccess uint32, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (uintptr, error) {
	service, err := windows.CreateService(
		windows.Handle(scManager),
		serviceName,
		databaseName,
		desiredAccess,
		serviceType,
		startType,
		errorControl,
		binaryPathName,
		loadOrderGroup,
		tagId,
		dependencies,
		serviceStartName,
		password,
	)
	if err != nil && err != windows.ERROR_SUCCESS {
		return 0, err
	}

	return uintptr(service), nil
}

func OpenServiceW(scManager uintptr, serviceName *uint16, desiredAccess uint32) (uintptr, error) {
	service, err := windows.OpenService(windows.Handle(scManager), serviceName, desiredAccess)
	if err != nil && err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return uintptr(service), nil
}

func StartServiceW(service uintptr, numServiceArgs uint32, serviceArgVectors **uint16) error {
	err := windows.StartService(windows.Handle(service), numServiceArgs, serviceArgVectors)
	if err != nil && err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func ControlService(service uintptr, control uint32, serviceStatus *windows.SERVICE_STATUS) error {
	err := windows.ControlService(windows.Handle(service), control, serviceStatus)
	if err != nil && err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func DeleteService(service uintptr) error {
	err := windows.DeleteService(windows.Handle(service))
	if err != nil && err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func CloseServiceHandle(scObject uintptr) error {
	err := windows.CloseServiceHandle(windows.Handle(scObject))
	if err != nil && err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
