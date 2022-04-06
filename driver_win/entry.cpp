#include "entry.h"


extern "C" DRIVER_INITIALIZE DriverEntry;

extern "C"
NTSTATUS
DriverEntry ( IN PDRIVER_OBJECT DriverObject, IN PUNICODE_STRING RegistryPath)
{


    
    return STATUS_SUCCESS;
}


