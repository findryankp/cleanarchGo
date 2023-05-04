package roots

func InitRoots(featuresName string) {
	dockerCreate(featuresName)
	gitigonreCreate(featuresName)
	envCreate(featuresName)
	loadConfigCreate(featuresName)
}
