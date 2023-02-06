package repository

import (
	"github.com/shirou/gopsutil/disk"
	"go-agent/entity"
)

type PartitionRepository struct {
}

func NewPartitionRepository() *PartitionRepository {
	return &PartitionRepository{}
}

func (PartitionRepository) GetPartitions() ([]*entity.Partition, error) {
	var patitionsEntity []*entity.Partition
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	for _, p := range partitions {
		patitionsEntity = append(patitionsEntity, &entity.Partition{
			Device:     p.Device,
			Mountpoint: p.Mountpoint,
			Fstype:     p.Fstype,
			Opts:       p.Opts,
		})
	}

	return patitionsEntity, nil
}
