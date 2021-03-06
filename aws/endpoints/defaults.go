// Code generated by aws/endpoints/v3model_codegen.go. DO NOT EDIT.

package endpoints

import (
	"regexp"
)

// Partition identifiers
const (
	Ks3PartitionID = "ks3" // KS3 partition.
)

// KS3 partition's regions.
const (
	Ks3CnBeijingRegionID          = "ks3-cn-beijing"           // Beijing.
	Ks3CnBeijingInternalRegionID  = "ks3-cn-beijing-internal"  // Beijing (internal).
	Ks3CnHongkongRegionID         = "ks3-cn-hongkong"          // Hongkong.
	Ks3CnHongkongInternalRegionID = "ks3-cn-hongkong-internal" // Hongkong (internal).
	Ks3CnShanghaiRegionID         = "ks3-cn-shanghai"          // Shanghai.
	Ks3CnShanghaiInternalRegionID = "ks3-cn-shanghai-internal" // Shanghai (internal).
	KssRegionID                   = "kss"                      // Hangzhou.
	KssInternalRegionID           = "kss-internal"             // Hangzhou (internal).
)

// Service identifiers
const (
	Ec2metadataServiceID = "ec2metadata" // Ec2metadata.
	S3ServiceID          = "s3"          // S3.
)

// DefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: KS3.
//
// Use DefaultPartitions() to get the list of the default partitions.
func DefaultResolver() Resolver {
	return defaultPartitions
}

// DefaultPartitions returns a list of the partitions the SDK is bundled
// with. The available partitions are: KS3.
//
//    partitions := endpoints.DefaultPartitions
//    for _, p := range partitions {
//        // ... inspect partitions
//    }
func DefaultPartitions() []Partition {
	return defaultPartitions.Partitions()
}

var defaultPartitions = partitions{
	ks3Partition,
}

// Ks3Partition returns the Resolver for KS3.
func Ks3Partition() Partition {
	return ks3Partition.Partition()
}

var ks3Partition = partition{
	ID:        "ks3",
	Name:      "KS3",
	DNSSuffix: "ksyun.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile(".*")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{region}.{dnsSuffix}",
		Protocols:         []string{"http", "https"},
		SignatureVersions: []string{"v2", "v4"},
	},
	Regions: regions{
		"ks3-cn-beijing": region{
			Description: "Beijing",
		},
		"ks3-cn-beijing-internal": region{
			Description: "Beijing (internal)",
		},
		"ks3-cn-hongkong": region{
			Description: "Hongkong",
		},
		"ks3-cn-hongkong-internal": region{
			Description: "Hongkong (internal)",
		},
		"ks3-cn-shanghai": region{
			Description: "Shanghai",
		},
		"ks3-cn-shanghai-internal": region{
			Description: "Shanghai (internal)",
		},
		"kss": region{
			Description: "Hangzhou",
		},
		"kss-internal": region{
			Description: "Hangzhou (internal)",
		},
	},
	Services: services{
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"s3": service{
			PartitionEndpoint: "kss",
			IsRegionalized:    boxedTrue,
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3", "s3v4"},
			},
			Endpoints: endpoints{
				"ks3-cn-beijing": endpoint{
					Hostname:          "ks3-cn-beijing.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ks3-cn-beijing-internal": endpoint{
					Hostname:          "ks3-cn-beijing-internal.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ks3-cn-hongkong": endpoint{
					Hostname:          "ks3-cn-hongkong.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ks3-cn-hongkong-internal": endpoint{
					Hostname:          "ks3-cn-hongkong-internal.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ks3-cn-shanghai": endpoint{
					Hostname:          "ks3-cn-shanghai.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ks3-cn-shanghai-internal": endpoint{
					Hostname:          "ks3-cn-shanghai-internal.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"kss": endpoint{
					Hostname:          "kss.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"kss-internal": endpoint{
					Hostname:          "kss-internal.ksyun.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
			},
		},
	},
}
