package util

type Flags interface {
	InitFlags() error
	PostProcessFlags() error
}

func InitAllFlags(slice ...Flags) error {
	for _, f := range slice {
		err := f.InitFlags()
		if err != nil {
			return err
		}
	}

	return nil
}

func PostProcessAllFlags(slice ...Flags) error {
	for _, f := range slice {
		err := f.PostProcessFlags()
		if err != nil {
			return err
		}
	}

	return nil
}
