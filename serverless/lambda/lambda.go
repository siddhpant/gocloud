package lambda

//GetFunction  describe lambda function.
func (lambda *Lambda) GetFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var getfunction Getfunction

	for key, value := range param {
		switch key {
		case "FunctionName":
			FunctionNameV, _ := value.(string)
			getfunction.FunctionName = FunctionNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Qualifier":
			QualifierV, _ := value.(string)
			getfunction.Qualifier = QualifierV
		}
	}

	params := make(map[string]string)
	preparegetfunctionqueryString(params, getfunction)

	response := make(map[string]interface{})
	err = Preparegetrequest(params, Region, response)

	resp = response
	return resp, err
}

//CreateFunction  Create lambda function.
func (lambda *Lambda) CreateFunction(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var region string
	var createfunction Createfunction

	for key, value := range param {
		switch key {

		case "Region":
			regionv, _ := value.(string)
			region = regionv

		case "FunctionName":
			FunctionNameV, _ := value.(string)
			createfunction.FunctionName = FunctionNameV

		case "Handler":
			handlerv, _ := value.(string)
			createfunction.handler = handlerv

		case "KMSKeyArn":
			kMSKeyArnv, _ := value.(string)
			createfunction.kMSKeyArn = kMSKeyArn

		case "MemorySize":
			memorySizev, _ := value.(int)
			createfunction.memorySizev = memorySizev

		case "Role":
			rolev, _ := value.(string)
			createfunction.role = rolev

		case "Publish":
			publishv, _ := value.(string)
			createfunction.publish = publishv

		case "Runtime":
			runtimev, _ := value.(string)
			createfunction.runtime = runtimev

		case "Tags":
			tagsv, _ := value.(string)
			createfunction.tags = tagsv

		case "Description":
			descriptionv, _ := value.(string)
			createfunction.description = descriptionv

		case "Timeout":
			timeoutv, _ := value.(int)
			createfunction.timeout = timeoutv

		case "DeadLetterConfig":
			deadLetterConfigparam, _ := value.(map[string]string)

			for deadLetterConfigparamkey, deadLetterConfigparamvalue := range deadLetterConfigparam {

				switch deadLetterConfigparam {
				case "TargetArn":
					targetArnv, _ := deadLetterConfigparamvalue.(string)
					createfunction.deadLetterConfig.targetArnv = targetArnv
				}

			}

		case "TracingConfig":
			tracingConfigparam, _ := value.(map[string]string)
			for tracingConfigparamkey, tracingConfigparamvalue := range tracingConfigparam {

				switch tracingConfigparam {
				case "Mode":
					modev, _ := tracingConfigparamvalue.(string)
					createfunction.tracingConfig.Mode = modev
				}

			}

		case "VpcConfig":
			vpcConfigparam, _ := value.(map[string]string)
			for vpcConfigparamkey, vpcConfigparamvalue := range vpcConfigparam {

				switch vpcConfigparam {
				case "SubnetIds":
					subnetIdsv, _ := vpcConfigvalue.([]string)
					createfunction.vpcConfig.subnetIds = subnetIdsv

				case "SecurityGroupIds":
					securityGroupIdsv, _ := vpcConfigvalue.([]string)
					createfunction.vpcConfig.securityGroupIds = securityGroupIds

				}
			}

		case "Code":
			codeparam, _ := value.(map[string]interface{})
			for codeparamkey, codeparamvalue := range codeparam {

				switch codeparamkey {
				case "S3Bucket":
					s3Bucketv, _ := codeparamvalue.(string)
					createfunction.code.s3Bucket = s3Bucketv

				case "S3Key":
					s3Keyv, _ := codeparamvalue.(string)
					createfunction.code.s3Key = s3Keyv

				case "S3ObjectVersion":
					s3ObjectVersionv, _ := codeparamvalue.(string)
					createfunction.code.s3ObjectVersion = s3ObjectVersionv

				case "ZipFile":
					zipFilev, _ := codeparamvalue.(string)
					createfunction.code.zipFile = zipFilev

				}
			}

		}
	}

	params := make(map[string]interface{})

	preparecreatefunctiondict(params, createfunction)

	response := make(map[string]interface{})

	err = PreparePostrequest(params, region, response)

	resp = response
	return resp, err
}

func preparecreatefunctiondict(params map[string]interface{}, createfunction Createfunction) {

	if createfunction.functionName != "" {
		params["FunctionName"] = createfunction.functionName
	}

	if createfunction.handler != "" {
		params["Handler"] = createfunction.handler
	}

	if createfunction.kMSKeyArn != "" {
		params["KMSKeyArn"] = createfunction.kMSKeyArn
	}

	if createfunction.memorySize > 0 {
		params["MemorySize"] = createfunction.memorySize
	}

	params["Publish"] = createfunction.publish

	if createfunction.role != 0 {
		params["Role"] = createfunction.role
	}

	if createfunction.role != "" {
		params["Role"] = createfunction.role
	}

	if createfunction.runtime != "" {
		params["Runtime"] = createfunction.runtime
	}

	if createfunction.tags != "" {
		param["Tags"] = createfunction.tags
	}

	if createfunction.description != "" {
		params["Description"] = createfunction.description
	}

	if createfunction.timeout > 0 {
		params["Timeout"] = createfunction.timeout
	}

	if createfunction.deadLetterConfig.targetArn != "" {
		param := request.(map[string]interface{})
		param["TargetArn"] = createfunction.deadLetterConfig.targetArn
		params["DeadLetterConfig"] = param
	}


		if createfunction.tracingConfig.mode != "" {
			param := request.(map[string]interface{})
			param["Mode"] = createfunction.tracingConfig.mode
			params["TracingConfig"] = param
		}

		if createfunction.tracingConfig.mode != "" {
			param := request.(map[string]interface{})
			param["Mode"] = createfunction.tracingConfig.mode
			params["TracingConfig"] = param
		}

		if (len(createfunction.vpcConfig.securityGroupIds > 0)  || (len(createfunction.vpcConfig.subnetIds) > 0) {
			param := request.(map[string]interface{})

			if(len(createfunction.vpcConfig.securityGroupIds > 0))
			{
				param["SecurityGroupIds"] = createfunction.vpcConfig.securityGroupIds
			}

			if(len(createfunction.vpcConfig.subnetIds) > 0)
			{
				param["SubnetIds"] = createfunction.vpcConfig.subnetIds
			}

			params["VpcConfig"] = param
		}

			preparecode(params ,createfunction)

}

func preparecode(params map[string]interface{}, createfunction Createfunction){

	param := request.(map[string]interface{})

	if createfunction.code.s3Bucket != "" {
		param["s3Bucket"] = createfunction.code.s3Bucket
	}

	if createfunction.code.s3Key != "" {
		param["Handler"] = createfunction.code.handler
	}

	if createfunction.code.s3ObjectVersion != "" {
		param["S3ObjectVersion"] = createfunction.code.s3ObjectVersion
	}


		if createfunction.code.zipFile != "" {
			param["ZipFile"] = createfunction.code.zipFile
		}

		params["Code"] = param

}

//CallFunction  call lambda function.
func (lambda *Lambda) CallFunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//ListFunction  list lambda function.
func (lambda *Lambda) ListFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var listfunction Listfunction

	for key, value := range param {
		switch key {

		case "FunctionVersion":
			functionVersionv, _ := value.(string)
			listfunction.functionVersion = functionVersionv

		case "Marker":
			markerv, _ := value.(string)
			listfunction.marker = markerv

		case "MasterRegion":
			masterRegionv, _ := value.(string)
			listfunction.masterRegion = masterRegionv

		case "MaxItems":
			maxItemsv, _ := value.(string)
			listfunction.maxItems = maxItemsv

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)
	preparelistfunctionqueryString(params, listfunction)

	response := make(map[string]interface{})
	err = Preparegetrequest(params, Region, response)

	resp = response
	return resp, err

}

//DeleteFunction  delete lambda function.
func (lambda *Lambda) DeleteFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var deletefunction Deletefunction

	for key, value := range param {
		switch key {
		case "FunctionName":
			FunctionNameV, _ := value.(string)
			deletefunction.FunctionName = FunctionNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Qualifier":
			QualifierV, _ := value.(string)
			deletefunction.Qualifier = QualifierV
		}
	}

	params := make(map[string]string)
	preparedeleteservicequeryString(params, deletefunction)

	response := make(map[string]interface{})
	err = Preparegetrequest(params, Region, response)

	resp = response
	return resp, err
}
