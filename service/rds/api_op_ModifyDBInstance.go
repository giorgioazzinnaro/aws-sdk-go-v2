// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies settings for a DB instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request. To learn what modifications you can make to your DB instance, call
// DescribeValidDBInstanceModifications before you call ModifyDBInstance.
func (c *Client) ModifyDBInstance(ctx context.Context, params *ModifyDBInstanceInput, optFns ...func(*Options)) (*ModifyDBInstanceOutput, error) {
	if params == nil {
		params = &ModifyDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBInstance", params, optFns, addOperationModifyDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBInstanceInput struct {

	// The DB instance identifier. This value is stored as a lowercase string.
	// Constraints:
	//
	//     * Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The new amount of storage (in gibibytes) to allocate for the DB instance. For
	// MariaDB, MySQL, Oracle, and PostgreSQL, the value supplied must be at least 10%
	// greater than the current value. Values that are not at least 10% greater than
	// the existing value are rounded up so that they are 10% greater than the current
	// value. For the valid values for allocated storage for each engine, see
	// CreateDBInstance.
	AllocatedStorage *int32

	// A value that indicates whether major version upgrades are allowed. Changing this
	// parameter doesn't result in an outage and the change is asynchronously applied
	// as soon as possible. Constraints: Major version upgrades must be allowed when
	// specifying a value for the EngineVersion parameter that is a different major
	// version than the DB instance's current version.
	AllowMajorVersionUpgrade *bool

	// A value that indicates whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB instance. By default, this
	// parameter is disabled. If this parameter is disabled, changes to the DB instance
	// are applied during the next maintenance window. Some parameter changes can cause
	// an outage and are applied on the next call to RebootDBInstance, or the next
	// failure reboot. Review the table of parameters in Modifying a DB Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	// in the Amazon RDS User Guide. to see the impact of enabling or disabling
	// ApplyImmediately for each modified parameter and to determine when the changes
	// are applied.
	ApplyImmediately *bool

	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window. Changing this parameter
	// doesn't result in an outage except in the following case and the change is
	// asynchronously applied as soon as possible. An outage results if this parameter
	// is enabled during the maintenance window, and a newer minor version is
	// available, and RDS has enabled auto patching for that engine version.
	AutoMinorVersionUpgrade *bool

	// The number of days to retain automated backups. Setting this parameter to a
	// positive number enables backups. Setting this parameter to 0 disables automated
	// backups. Changing this parameter can result in an outage if you change from 0 to
	// a non-zero value or from a non-zero value to 0. These changes are applied during
	// the next maintenance window unless the ApplyImmediately parameter is enabled for
	// this request. If you change the parameter from one non-zero value to another
	// non-zero value, the change is asynchronously applied as soon as possible. Amazon
	// Aurora Not applicable. The retention period for automated backups is managed by
	// the DB cluster. For more information, see ModifyDBCluster. Default: Uses
	// existing setting Constraints:
	//
	//     * Must be a value from 0 to 35
	//
	//     * Can be
	// specified for a MySQL read replica only if the source is running MySQL 5.6 or
	// later
	//
	//     * Can be specified for a PostgreSQL read replica only if the source
	// is running PostgreSQL 9.3.5
	//
	//     * Can't be set to 0 if the DB instance is a
	// source to read replicas
	BackupRetentionPeriod *int32

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string

	// A value that indicates whether the DB instance is restarted when you rotate your
	// SSL/TLS certificate. By default, the DB instance is restarted when you rotate
	// your SSL/TLS certificate. The certificate is not updated until the DB instance
	// is restarted. Set this parameter only if you are not using SSL/TLS to connect to
	// the DB instance. If you are using SSL/TLS to connect to the DB instance, follow
	// the appropriate instructions for your DB engine to rotate your SSL/TLS
	// certificate:
	//
	//     * For more information about rotating your SSL/TLS certificate
	// for RDS DB engines, see  Rotating Your SSL/TLS Certificate.
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html)
	// in the Amazon RDS User Guide.
	//
	//     * For more information about rotating your
	// SSL/TLS certificate for Aurora DB engines, see  Rotating Your SSL/TLS
	// Certificate
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html)
	// in the Amazon Aurora User Guide.
	CertificateRotationRestart *bool

	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB instance. A change to the
	// CloudwatchLogsExportConfiguration parameter is always applied to the DB instance
	// immediately. Therefore, the ApplyImmediately parameter has no effect.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// A value that indicates whether to copy all tags from the DB instance to
	// snapshots of the DB instance. By default, tags are not copied. Amazon Aurora Not
	// applicable. Copying tags to snapshots is managed by the DB cluster. Setting this
	// value for an Aurora DB instance has no effect on the DB cluster setting. For
	// more information, see ModifyDBCluster.
	CopyTagsToSnapshot *bool

	// The new compute and memory capacity of the DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all AWS Regions, or
	// for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. If you modify the DB instance class, an outage
	// occurs during the change. The change is applied during the next maintenance
	// window, unless ApplyImmediately is enabled for this request. Default: Uses
	// existing setting
	DBInstanceClass *string

	// The name of the DB parameter group to apply to the DB instance. Changing this
	// setting doesn't result in an outage. The parameter group name itself is changed
	// immediately, but the actual parameter changes are not applied until you reboot
	// the instance without failover. In this case, the DB instance isn't rebooted
	// automatically and the parameter changes isn't applied during the next
	// maintenance window. Default: Uses existing setting Constraints: The DB parameter
	// group must be in the same DB parameter group family as this DB instance.
	DBParameterGroupName *string

	// The port number on which the database accepts connections. The value of the
	// DBPortNumber parameter must not match any of the port values specified for
	// options in the option group for the DB instance. Your database will restart when
	// you change the DBPortNumber value regardless of the value of the
	// ApplyImmediately parameter. MySQL Default: 3306 Valid values: 1150-65535 MariaDB
	// Default: 3306 Valid values: 1150-65535 PostgreSQL Default: 5432 Valid values:
	// 1150-65535 Type: Integer Oracle Default: 1521 Valid values: 1150-65535 SQL
	// Server Default: 1433 Valid values: 1150-65535 except 1234, 1434, 3260, 3343,
	// 3389, 47001, and 49152-49156. Amazon Aurora Default: 3306 Valid values:
	// 1150-65535
	DBPortNumber *int32

	// A list of DB security groups to authorize on this DB instance. Changing this
	// setting doesn't result in an outage and the change is asynchronously applied as
	// soon as possible. Constraints:
	//
	//     * If supplied, must match existing
	// DBSecurityGroups.
	DBSecurityGroups []*string

	// The new DB subnet group for the DB instance. You can use this parameter to move
	// your DB instance to a different VPC. If your DB instance isn't in a VPC, you can
	// also use this parameter to move your DB instance into a VPC. For more
	// information, see Updating the VPC for a DB Instance
	// (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html#USER_VPC.Non-VPC2VPC)
	// in the Amazon RDS User Guide. Changing the subnet group causes an outage during
	// the change. The change is applied during the next maintenance window, unless you
	// enable ApplyImmediately. Constraints: If supplied, must match the name of an
	// existing DBSubnetGroup. Example: mySubnetGroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see  Deleting a DB
	// Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool

	// The Active Directory directory ID to move the DB instance to. Specify none to
	// remove the instance from its current domain. The domain must be created prior to
	// this operation. Currently, only MySQL, Microsoft SQL Server, Oracle, and
	// PostgreSQL DB instances can be created in an Active Directory Domain. For more
	// information, see  Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide.
	Domain *string

	// The name of the IAM role to use when making API calls to the Directory Service.
	DomainIAMRoleName *string

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance. For more
	// information about IAM database authentication, see  IAM Database Authentication
	// for MySQL and PostgreSQL
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool

	// A value that indicates whether to enable Performance Insights for the DB
	// instance. For more information, see Using Amazon Performance Insights
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon Relational Database Service User Guide.
	EnablePerformanceInsights *bool

	// The version number of the database engine to upgrade to. Changing this parameter
	// results in an outage and the change is applied during the next maintenance
	// window unless the ApplyImmediately parameter is eanbled for this request. For
	// major version upgrades, if a nondefault DB parameter group is currently in use,
	// a new DB parameter group in the DB parameter group family for the new engine
	// version must be specified. The new DB parameter group can be the default for
	// that DB parameter group family. For information about valid engine versions, see
	// CreateDBInstance, or call DescribeDBEngineVersions.
	EngineVersion *string

	// The new Provisioned IOPS (I/O operations per second) value for the RDS instance.
	// Changing this setting doesn't result in an outage and the change is applied
	// during the next maintenance window unless the ApplyImmediately parameter is
	// enabled for this request. If you are migrating from Provisioned IOPS to standard
	// storage, set this value to 0. The DB instance will require a reboot for the
	// change in storage type to take effect. If you choose to migrate your DB instance
	// from using standard storage to using Provisioned IOPS, or from using Provisioned
	// IOPS to using standard storage, the process can take time. The duration of the
	// migration depends on several factors such as database load, storage size,
	// storage type (standard or Provisioned IOPS), amount of IOPS provisioned (if
	// any), and the number of prior scale storage operations. Typical migration times
	// are under 24 hours, but the process can take up to several days in some cases.
	// During the migration, the DB instance is available for use, but might experience
	// performance degradation. While the migration takes place, nightly backups for
	// the instance are suspended. No other Amazon RDS operations can take place for
	// the instance, including modifying the instance, rebooting the instance, deleting
	// the instance, creating a read replica for the instance, and creating a DB
	// snapshot of the instance. Constraints: For MariaDB, MySQL, Oracle, and
	// PostgreSQL, the value supplied must be at least 10% greater than the current
	// value. Values that are not at least 10% greater than the existing value are
	// rounded up so that they are 10% greater than the current value. Default: Uses
	// existing setting
	Iops *int32

	// The license model for the DB instance. Valid values: license-included |
	// bring-your-own-license | general-public-license
	LicenseModel *string

	// The new password for the master user. The password can include any printable
	// ASCII character except "/", """, or "@". Changing this parameter doesn't result
	// in an outage and the change is asynchronously applied as soon as possible.
	// Between the time of the request and the completion of the request, the
	// MasterUserPassword element exists in the PendingModifiedValues element of the
	// operation response. Amazon Aurora Not applicable. The password for the master
	// user is managed by the DB cluster. For more information, see ModifyDBCluster.
	// Default: Uses existing setting MariaDB Constraints: Must contain from 8 to 41
	// characters. Microsoft SQL Server Constraints: Must contain from 8 to 128
	// characters. MySQL Constraints: Must contain from 8 to 41 characters. Oracle
	// Constraints: Must contain from 8 to 30 characters. PostgreSQL Constraints: Must
	// contain from 8 to 128 characters. Amazon RDS API actions never return the
	// password, so this action provides a way to regain access to a primary instance
	// user if the password is lost. This includes restoring privileges that might have
	// been accidentally revoked.
	MasterUserPassword *string

	// The upper limit to which Amazon RDS can automatically scale the storage of the
	// DB instance.
	MaxAllocatedStorage *int32

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, then
	// you must also set MonitoringInterval to a value other than 0. Valid Values: 0,
	// 1, 5, 10, 15, 30, 60
	MonitoringInterval *int32

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to
	// Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For
	// information on creating a monitoring role, go to To create an IAM role for
	// Amazon RDS Enhanced Monitoring
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html#USER_Monitoring.OS.IAMRole)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than
	// 0, then you must supply a MonitoringRoleArn value.
	MonitoringRoleArn *string

	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	// Changing this parameter doesn't result in an outage and the change is applied
	// during the next maintenance window unless the ApplyImmediately parameter is
	// enabled for this request.
	MultiAZ *bool

	// The new DB instance identifier for the DB instance when renaming a DB instance.
	// When you change the DB instance identifier, an instance reboot occurs
	// immediately if you enable ApplyImmediately, or will occur during the next
	// maintenance window if you disable Apply Immediately. This value is stored as a
	// lowercase string. Constraints:
	//
	//     * Must contain from 1 to 63 letters,
	// numbers, or hyphens.
	//
	//     * The first character must be a letter.
	//
	//     * Can't
	// end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	NewDBInstanceIdentifier *string

	// Indicates that the DB instance should be associated with the specified option
	// group. Changing this parameter doesn't result in an outage except in the
	// following case and the change is applied during the next maintenance window
	// unless the ApplyImmediately parameter is enabled for this request. If the
	// parameter change results in an option group that enables OEM, this change can
	// cause a brief (sub-second) period during which new connections are rejected but
	// existing connections are not interrupted. Permanent options, such as the TDE
	// option for Oracle Advanced Security TDE, can't be removed from an option group,
	// and that option group can't be removed from a DB instance once it is associated
	// with a DB instance
	OptionGroupName *string

	// The AWS KMS key identifier for encryption of Performance Insights data. The KMS
	// key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key
	// alias for the KMS encryption key. If you do not specify a value for
	// PerformanceInsightsKMSKeyId, then Amazon RDS uses your default encryption key.
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	PerformanceInsightsKMSKeyId *string

	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, as determined by the BackupRetentionPeriod parameter.
	// Changing this parameter doesn't result in an outage and the change is
	// asynchronously applied as soon as possible. Amazon Aurora Not applicable. The
	// daily time range for creating automated backups is managed by the DB cluster.
	// For more information, see ModifyDBCluster. Constraints:
	//
	//     * Must be in the
	// format hh24:mi-hh24:mi
	//
	//     * Must be in Universal Time Coordinated (UTC)
	//
	//     *
	// Must not conflict with the preferred maintenance window
	//
	//     * Must be at least
	// 30 minutes
	PreferredBackupWindow *string

	// The weekly time range (in UTC) during which system maintenance can occur, which
	// might result in an outage. Changing this parameter doesn't result in an outage,
	// except in the following situation, and the change is asynchronously applied as
	// soon as possible. If there are pending actions that cause a reboot, and the
	// maintenance window is changed to include the current time, then changing this
	// parameter will cause a reboot of the DB instance. If moving this window to the
	// current time, there must be at least 30 minutes between the current time and end
	// of the window to ensure pending changes are applied. Default: Uses existing
	// setting Format: ddd:hh24:mi-ddd:hh24:mi Valid Days: Mon | Tue | Wed | Thu | Fri
	// | Sat | Sun Constraints: Must be at least 30 minutes
	PreferredMaintenanceWindow *string

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []*types.ProcessorFeature

	// A value that specifies the order in which an Aurora Replica is promoted to the
	// primary instance after a failure of the existing primary instance. For more
	// information, see  Fault Tolerance for an Aurora DB Cluster
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance)
	// in the Amazon Aurora User Guide. Default: 1 Valid Values: 0 - 15
	PromotionTier *int32

	// A value that indicates whether the DB instance is publicly accessible. When the
	// DB instance is publicly accessible, its DNS endpoint resolves to the private IP
	// address from within the DB instance's VPC, and to the public IP address from
	// outside of the DB instance's VPC. Access to the DB instance is ultimately
	// controlled by the security group it uses, and that public access is not
	// permitted if the security group assigned to the DB instance doesn't permit it.
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address. PubliclyAccessible only
	// applies to DB instances in a VPC. The DB instance must be part of a public
	// subnet and PubliclyAccessible must be enabled for it to be publicly accessible.
	// Changes to the PubliclyAccessible parameter are applied immediately regardless
	// of the value of the ApplyImmediately parameter.
	PubliclyAccessible *bool

	// A value that sets the open mode of a replica database to either mounted or
	// read-only. Currently, this parameter is only supported for Oracle DB instances.
	// Mounted DB replicas are included in Oracle Enterprise Edition. The main use case
	// for mounted replicas is cross-Region disaster recovery. The primary database
	// doesn't use Active Data Guard to transmit information to the mounted replica.
	// Because it doesn't accept user connections, a mounted replica can't serve a
	// read-only workload. For more information, see Working with Oracle Read Replicas
	// for Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)
	// in the Amazon RDS User Guide.
	ReplicaMode types.ReplicaMode

	// Specifies the storage type to be associated with the DB instance. If you specify
	// Provisioned IOPS (io1), you must also include a value for the Iops parameter. If
	// you choose to migrate your DB instance from using standard storage to using
	// Provisioned IOPS, or from using Provisioned IOPS to using standard storage, the
	// process can take time. The duration of the migration depends on several factors
	// such as database load, storage size, storage type (standard or Provisioned
	// IOPS), amount of IOPS provisioned (if any), and the number of prior scale
	// storage operations. Typical migration times are under 24 hours, but the process
	// can take up to several days in some cases. During the migration, the DB instance
	// is available for use, but might experience performance degradation. While the
	// migration takes place, nightly backups for the instance are suspended. No other
	// Amazon RDS operations can take place for the instance, including modifying the
	// instance, rebooting the instance, deleting the instance, creating a read replica
	// for the instance, and creating a DB snapshot of the instance. Valid values:
	// standard | gp2 | io1 Default: io1 if the Iops parameter is specified, otherwise
	// gp2
	StorageType *string

	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string

	// A value that indicates whether the DB instance class of the DB instance uses its
	// default processor features.
	UseDefaultProcessorFeatures *bool

	// A list of EC2 VPC security groups to authorize on this DB instance. This change
	// is asynchronously applied as soon as possible. Amazon Aurora Not applicable. The
	// associated list of EC2 VPC security groups is managed by the DB cluster. For
	// more information, see ModifyDBCluster. Constraints:
	//
	//     * If supplied, must
	// match existing VpcSecurityGroupIds.
	VpcSecurityGroupIds []*string
}

type ModifyDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(middleware.After, &awsAwsquery_serializeOpModifyDBInstance{})
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(middleware.After, &awsAwsquery_deserializeOpModifyDBInstance{})
	if err != nil {
		return err
	}
	if err := awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err := addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err := v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err := addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err := addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err := awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err := addClientUserAgent(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err := smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err := addOpModifyDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err := stack.Initialize.Add(middleware.Before, newServiceMetadataMiddleware_opModifyDBInstance(options.Region)); err != nil {
		return err
	}
	if err := addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err := addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBInstance",
	}
}
