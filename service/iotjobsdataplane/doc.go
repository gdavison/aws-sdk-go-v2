// Code generated by smithy-go-codegen DO NOT EDIT.

// Package iotjobsdataplane provides the API client, operations, and parameter
// types for AWS IoT Jobs Data Plane.
//
// AWS IoT Jobs is a service that allows you to define a set of jobs — remote
// operations that are sent to and executed on one or more devices connected to AWS
// IoT. For example, you can define a job that instructs a set of devices to
// download and install application or firmware updates, reboot, rotate
// certificates, or perform remote troubleshooting operations.
//
// To create a job, you make a job document which is a description of the remote
// operations to be performed, and you specify a list of targets that should
// perform the operations. The targets can be individual things, thing groups or
// both.
//
// AWS IoT Jobs sends a message to inform the targets that a job is available. The
// target starts the execution of the job by downloading the job document,
// performing the operations it specifies, and reporting its progress to AWS IoT.
// The Jobs service provides commands to track the progress of a job on a specific
// target and for all the targets of the job
package iotjobsdataplane
