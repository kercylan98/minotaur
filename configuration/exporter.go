package configuration

import "github.com/kercylan98/minotaur/configuration/raw"

// Exporter 数据导出器
type Exporter interface {
	// Export 导出数据
	Export(config raw.Config, data any) error
}

// ExportData 导出数据
func ExportData(exporter Exporter, scanners ...Scanner) error {
	for _, scanner := range scanners {
		config, err := scanner.StructScan()
		if err != nil {
			return err
		}

		data, err := scanner.DataScan(config.GetFieldsWithSort())
		if err != nil {
			return err
		}
		if err = exporter.Export(config, data); err != nil {
			return err
		}
	}
	return nil
}
