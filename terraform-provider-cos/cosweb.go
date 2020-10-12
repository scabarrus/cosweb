package cos

import(
	"log"
	"net/http"
	"crypto/tls"
	//"strings"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

//IBMToken struct store all attributes of the token provided by IBM IAM endpoint.
type IBMToken struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ImsUserID string `json:"ims_user_id"`
	TokenType string `json:"token_type"`
	ExpireIn int `json:"expires_in"`
	Expiration int `json:"expiration"`
	RefreshTokenExpiration int `json:"refresh_token_expiration"`
	Scope string `json:"scope"` 
}

type CosWeb struct{
	ApiKey string
	BearerToken string
	Endpoint string
	Protocol string
	SslCheck bool
}

//CosInstanceService is a JSON struct that match with cos instance object
type CosInstance struct{
	Name string `json:"name"`
	Description string `json:"description"`
	ResourceGroup string `json:"resourceGroup"`
	ResourcePlanId string `json:"resourcePlanId"`
	ID int `json:"id"`
	//CreateAt string `json:-"`
}

//Bucket is a bucket structure
type Bucket struct{
	Name string `json:"name"`
	Description string `json:"description"`
	CosInstanceGUID uint `json:"cosInstanceGUID"`
}

func (c CosWeb) Initialize(apiKey string, bearer string,endpoint string, protocol string, sslcheck bool)(* CosWeb){
	return &CosWeb{apiKey,bearer,endpoint,protocol,sslcheck}
}

//CallRest function make a REST call 
//It formats the request according to data passed
//It return the http response and/or an error
func(c CosWeb) CallRest(url string,headers map[string]string,method string,payload *bytes.Buffer)(*http.Response,error){
	if c.SslCheck == false {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	log.Printf("[DEBUG] URL :",url)
	log.Printf("[DEBUG] METHOD: %s",method)
	log.Printf("[DEBUG] HEADER : %s",headers)
	log.Println("[DEBUG] PAYLOAD : %s",payload)
	request, _ := http.NewRequest(method, url,payload)
	for header,value:= range headers {
		request.Header.Add(header, value)
	}
	response, err := http.DefaultClient.Do(request)
	log.Println(err)
	return response,err
}



//GetBucket method call cosweb microservice to get a bucket from cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) GetAllCosInstances()([]CosInstance,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances"
	method:="GET"
	var headers =make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	payload :=  bytes.Buffer{}
	var cos []CosInstance
	response,err:=c.CallRest(url,headers,method,&payload)
	if err != nil{
		return []CosInstance{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		log.Printf("body : %s",response.Body)
	
	_ = json.Unmarshal(body, &cos)
	}
	return cos,err
}
//CreateCosInstance method call cosweb microservice to create a cos instance on IBM cloud and register it in cosweb database
//It expects the cos instance name, description, resource group and resource plan id as parameters
//It returns CosInstance struct or an error
func (c CosWeb)CreateCosInstance(cosInstanceName string,cosInstanceDescription string,cosInstanceResourceGroup string,cosInstanceResourcePlanId string)(CosInstance,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances"
	method:="POST"
	headers:=make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	var instance CosInstance = CosInstance{cosInstanceName,cosInstanceDescription,cosInstanceResourceGroup,cosInstanceResourcePlanId,0}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(instance)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return CosInstance{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &instance)
	log.Printf("body : ",instance)
	}
	return instance,err
}

//GetCosInstance method call cosweb microservice to retrieve a cos instance on IBM cloud and register it in cosweb database
//It expects the cos instance id as parameters
//It returns CosInstance struct or an error
func (c CosWeb)GetCosInstance(cosInstanceId int)(CosInstance,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)
	method:="GET"
	headers:=make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	var instance CosInstance = CosInstance{}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(instance)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return CosInstance{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &instance)
	log.Printf("body : ",instance)
	}
	return instance,err
}

//ModifyCosInstance method call cosweb microservice to modify a cos instance on IBM cloud and register it in cosweb database
//It expects the cos instance id, cos instance name, description, resource group and resource plan id as parameters
//It returns CosInstance struct or an error
func (c CosWeb)ModifyCosInstance(cosInstanceId int,cosInstanceName string,cosInstanceDescription string,cosInstanceResourceGroup string,cosInstanceResourcePlanId string)(CosInstance,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)
	method:="PUT"
	headers:=make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	var instance CosInstance = CosInstance{cosInstanceName,cosInstanceDescription,cosInstanceResourceGroup,cosInstanceResourcePlanId,0}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(instance)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return CosInstance{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &instance)
	log.Printf("body : ",instance)
	}
	return instance,err
}

//DeleteCosInstance method call cosweb microservice to delete a cos instance on IBM cloud and register it in cosweb database
//It expects the cos instance id, cos instance name, description, resource group and resource plan id as parameters
//It returns CosInstance struct or an error
func (c CosWeb)DeleteCosInstance(cosInstanceId int,cosInstanceName string,cosInstanceDescription string,cosInstanceResourceGroup string,cosInstanceResourcePlanId string)(CosInstance,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)
	method:="DELETE"
	headers:=make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	var instance CosInstance = CosInstance{cosInstanceName,cosInstanceDescription,cosInstanceResourceGroup,cosInstanceResourcePlanId,0}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(instance)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return CosInstance{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &instance)
	log.Printf("body : ",instance)
	}
	return instance,err
}

//CreateBucket method call cosweb microservice to create a bucket on IBM cloud and register it in cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) CreateBucket(cosInstanceId int, bucketName string, bucketDescription string)(Bucket,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)+"/buckets"
	method:="POST"
	headers:= make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken

	
	var b Bucket = Bucket{bucketName,bucketDescription,uint(cosInstanceId)}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(b)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return Bucket{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &b)
	log.Printf("body : ",b)
	}
	return b,err
}

//GetBucket method call cosweb microservice to get a bucket from cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) GetAllBuckets(cosInstanceId int)([]Bucket,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)+"/buckets"
	method:="GET"
	var headers =make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	payload :=  bytes.Buffer{}
	var b []Bucket
	response,err:=c.CallRest(url,headers,method,&payload)
	if err != nil{
		return []Bucket{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		log.Printf("body : ",response.Body)
	
	_ = json.Unmarshal(body, &b)
	}
	return b,err
}

//GetBucket method call cosweb microservice to get a bucket from cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) GetBucket(cosInstanceId int,bucketId int)(Bucket,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)+"/buckets/"+strconv.Itoa(bucketId)
	method:="GET"
	var headers =make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken
	payload :=  bytes.Buffer{}
	var b Bucket
	response,err:=c.CallRest(url,headers,method,&payload)
	if err != nil{
		return Bucket{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		log.Printf("body : ",response.Body)
	
	_ = json.Unmarshal(body, &b)
	}
	return b,err
}

//ModifyBucket method call cosweb microservice to modify a bucket on IBM cloud and register it in cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) ModifyBucket(cosInstanceId int,bucketId int, bucketName string, bucketDescription string)(Bucket,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)+"/buckets/"+strconv.Itoa(bucketId)
	method:="PUT"
	headers:= make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken

	
	var b Bucket = Bucket{bucketName,bucketDescription,uint(cosInstanceId)}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(b)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return Bucket{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &b)
	log.Printf("body : ",b)
	}
	return b,err
}

//DeleteBucket method call cosweb microservice to delete a bucket on IBM cloud and register it in cosweb database
//It expects apikey, the cos instance id, bucket name and bucket description as parameters
//It returns Bucket struct or an error
func (c CosWeb) DeleteBucket(cosInstanceId int,bucketId int, bucketName string, bucketDescription string)(Bucket,error){
	url:=c.Protocol+"://"+c.Endpoint+"/cosinstances/"+strconv.Itoa(cosInstanceId)+"/buckets/"+strconv.Itoa(bucketId)
	method:="DELETE"
	headers:= make(map[string]string)
	headers["x-api-key"]=c.ApiKey
	headers["Authorization"]="Bearer "+c.BearerToken

	
	var b Bucket = Bucket{bucketName,bucketDescription,uint(cosInstanceId)}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(b)
	response,err:=c.CallRest(url,headers,method,payload)
	if err != nil{
		return Bucket{},err
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//IBM cloud sent a success return code
	if response.StatusCode >= 200 && response.StatusCode < 300{
		
		log.Printf("status : ",response.StatusCode)
	
	_ = json.Unmarshal(body, &b)
	log.Printf("body : ",b)
	}
	return b,err
}



