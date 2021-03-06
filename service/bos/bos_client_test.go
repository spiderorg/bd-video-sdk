package bos

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spiderorg/bd-video-sdk/auth"
)

const (
	DefaultAccessKeyId     = "aaaaaaaaaaaaaaaa"
	DefaultSecretAccessKey = "bbbbbbbbbbbbbbb"
	DefaultDebugHost       = "bos.bj.baidubce.com"
	TestBukketName         = "baidubce-golang-sdk-test"
	TestObjectName         = "baidubce-golang-sdk-test-object"
	TestObjectName1        = "baidubce-golang-sdk-test-object-1"
)

var AccessKeyId string
var SecretAccessKey string
var DebugHost string

func TestInit(t *testing.T) {
	AccessKeyId = DefaultAccessKeyId
	if os.Getenv("ACCESS_KEY_ID") != "" {
		AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	}

	SecretAccessKey = DefaultSecretAccessKey
	if os.Getenv("SECRET_ACCESS_KEY") != "" {
		SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	}

	DebugHost = DefaultDebugHost
	if os.Getenv("DEBUG_HOST") != "" {
		DebugHost = os.Getenv("DEBUG_HOST")
	}

	fileSize := 256
	content := make([]byte, fileSize)
	rand.Read(content)
	err := ioutil.WriteFile(TestObjectName, content, 0644)
	if err != nil {
		t.Errorf("Write TestFile failed.")
	}

	fileSize = 1024*1024*5 + 1
	content = make([]byte, fileSize)
	rand.Read(content)
	err = ioutil.WriteFile(TestObjectName1, content, 0644)
	if err != nil {
		t.Errorf("Write TestFile failed.")
	}

}

func TestNewBosClient(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}

	if c.GetEndpoint() != "http://bos.bj.baidubce.com" {
		t.Errorf("GetEndpoint failed.")
	}

	if c.GetBaseURL() != "http://bos.bj.baidubce.com/v1" {
		t.Errorf("GetBaseURL failed.")
	}

	c.Location = "gz"
	c.APIVersion = "v2"

	if c.GetEndpoint() != "http://bos.gz.baidubce.com" {
		t.Errorf("GetEndpoint failed.")
	}

	if c.GetBaseURL() != "http://bos.gz.baidubce.com/v2" {
		t.Errorf("GetBaseURL failed.")
	}

	c.Host = "www.baidu.com"
	if c.GetEndpoint() != "http://www.baidu.com" {
		t.Errorf("GetEndpoint failed.")
	}

	if c.GetBaseURL() != "http://www.baidu.com/v2" {
		t.Errorf("GetBaseURL failed.")
	}

}

func TestPutBucket(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	if c.HeadBucket(TestBukketName) == nil {
	}
	c.DeleteBucket(TestBukketName)
	err = c.PutBucket(TestBukketName)
	if err != nil {
		t.Errorf("PutBucket failed.")
		t.Errorf(err.Error())
	}
}
func TestListBucket(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	_, err = c.ListBucket()
	if err != nil {
		t.Errorf("ListBucket failed.")
		t.Errorf(err.Error())
	}
}

func TestGetBucketAcl(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	_, err = c.GetBucketAcl(TestBukketName)
	if err != nil {
		t.Errorf("GetBucketAcl failed.")
		t.Errorf(err.Error())
	}
}

func TestListObjects(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	_, err = c.ListObjects(TestBukketName, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ListObjects failed.")
		t.Errorf(err.Error())
	}
}

func TestHeadBucket(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	err = c.HeadBucket(TestBukketName)
	if err != nil {
		t.Errorf("HeadBucket failed.")
		t.Errorf(err.Error())
	}
}

func TestGetBucketLocation(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	_, err = c.GetBucketLocation(TestBukketName)
	if err != nil {
		t.Errorf("GetBucketLocation failed.")
		t.Errorf(err.Error())
	}
}

func TestSetBucketAcl(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	err = c.SetBucketAcl(TestBukketName, "private")
	if err != nil {
		t.Errorf("SetBucketAcl failed.")
		t.Errorf(err.Error())
	}
}

func TestDeleteBucket(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	if c.HeadBucket(TestBukketName) != nil {
		c.PutBucket(TestBukketName)
	}
	err = c.DeleteBucket(TestBukketName)
	if err != nil {
		t.Errorf("DeleteBucket failed.")
		t.Errorf(err.Error())
	}
}

func TestPutObject(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	file, err := os.Open(TestObjectName)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	err = c.PutBucket(TestBukketName)
	eTag, err := c.PutObject(TestBukketName, TestObjectName, bytes.NewReader(content), "", "", map[string]string{"TEST": "2333"})
	if err != nil {
		t.Errorf("PutObject failed.")
		t.Errorf(err.Error())
	}
	h := md5.New()
	io.WriteString(h, string(content))
	if eTag != fmt.Sprintf("%x", h.Sum(nil)) {
		t.Errorf("PutObject failed. eTag Not Match.")
	}
}

func TestGetObjectMeta(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.GetObjectMeta(TestBukketName, TestObjectName)
	if err != nil {
		t.Errorf("GetObjectMeta failed.")
		t.Errorf(err.Error())
	}
	if res["eTag"] == "" {
		t.Errorf("GetObjectMeta failed. eTag Not Match")
	}
}

func TestDeleteObject(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	err = c.DeleteObject(TestBukketName, TestObjectName)
	if err != nil {
		t.Errorf("DeleteObject failed.")
		t.Errorf(err.Error())
	}
}

func TestInitiateMultipartUpload(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.InitiateMultipartUpload(TestBukketName, TestObjectName, "")
	if err != nil {
		t.Errorf("InitiateMultipartUpload failed.")
		t.Errorf(err.Error())
	}
	if res.BucketName != TestBukketName {
		t.Errorf("InitiateMultipartUpload failed. Bucket dont Match")
	}
	if res.ObjectName != TestObjectName {
		t.Errorf("InitiateMultipartUpload failed. Key dont Match")
	}
	if res.UploadId == "" {
		t.Errorf("InitiateMultipartUpload failed. UploadId dont Match")
	}
}

func TestListMultipartUploads(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.ListMultipartUploads(TestBukketName, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ListMultipartUploads failed.")
		t.Errorf(err.Error())
	}
	if res.BucketName != TestBukketName {
		t.Errorf("ListMultipartUploads failed. TestBukketName not Match")
	}
}

func TestCompleteMultipartUpload(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.InitiateMultipartUpload(TestBukketName, TestObjectName, "")
	uploadId := res.UploadId

	file, err := os.Open(TestObjectName1)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	eTag, err := c.UploadPart(TestBukketName, TestObjectName, uploadId, "1", bytes.NewReader(content))
	if err != nil {
		t.Errorf("UploadPart failed.")
	}
	partInfo1 := PartInfo{PartNumber: 1, ETag: eTag}

	file1, err := os.Open(TestObjectName)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file1.Close()
	content1, err := ioutil.ReadAll(file1)

	eTag, err = c.UploadPart(TestBukketName, TestObjectName, uploadId, "2", bytes.NewReader(content1))
	if err != nil {
		t.Errorf("UploadPart failed.")
	}
	partInfo2 := PartInfo{PartNumber: 2, ETag: eTag}

	parts := []PartInfo{partInfo1, partInfo2}

	_, err = c.CompleteMultipartUpload(TestBukketName, TestObjectName, uploadId, parts)
	if err != nil {
		t.Errorf("CompleteMultipartUpload failed.")
		t.Errorf(err.Error())
	}
	c.DeleteObject(TestBukketName, TestObjectName)
}

func TestAbortMultipartUpload(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.InitiateMultipartUpload(TestBukketName, TestObjectName, "")
	uploadId := res.UploadId

	err = c.AbortMultipartUpload(TestBukketName, TestObjectName, uploadId)
	if err != nil {
		t.Errorf("AbortMultipartUpload failed.")
	}
}

func TestListParts(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	res, err := c.InitiateMultipartUpload(TestBukketName, TestObjectName, "")
	uploadId := res.UploadId

	file, err := os.Open(TestObjectName1)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	eTag1, err := c.UploadPart(TestBukketName, TestObjectName, uploadId, "1", bytes.NewReader(content))
	if err != nil {
		t.Errorf("UploadPart failed.")
	}

	file1, err := os.Open(TestObjectName)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file1.Close()
	content1, err := ioutil.ReadAll(file1)

	eTag2, err := c.UploadPart(TestBukketName, TestObjectName, uploadId, "2", bytes.NewReader(content1))
	if err != nil {
		t.Errorf("UploadPart failed.")
	}

	res1, err := c.ListParts(TestBukketName, TestObjectName, uploadId, nil, nil)
	if err != nil {
		t.Errorf("ListParts failed.")
		t.Errorf(err.Error())
	}
	if res1.Parts[0].ETag != eTag1 {
		t.Errorf("ListParts failed. eTag1 Not Match")
	}
	if res1.Parts[1].ETag != eTag2 {
		t.Errorf("ListParts failed. eTag2 Not Match")
	}
}

func TestCopyObject(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	file, err := os.Open(TestObjectName)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	err = c.PutBucket(TestBukketName)
	_, err = c.PutObject(TestBukketName, TestObjectName, bytes.NewReader(content), "", "", nil)
	if err != nil {
		t.Errorf("PutObject failed.")
		t.Errorf(err.Error())
	}

	res1, err := c.CopyObject(TestBukketName, TestObjectName, TestBukketName, TestObjectName1, "", "")
	if err != nil {
		t.Errorf("CopyObject failed.")
		t.Errorf(err.Error())
	}

	h := md5.New()
	io.WriteString(h, string(content))

	if res1.ETag != fmt.Sprintf("%x", h.Sum(nil)) {
		t.Errorf("CopyObject failed. eTag Not Match.")
	}

	c.DeleteObject(TestBukketName, TestObjectName)
	c.DeleteObject(TestBukketName, TestObjectName1)
}

func TestGetObject(t *testing.T) {
	c, err := NewBosClient(auth.NewBceCredentials(AccessKeyId, SecretAccessKey))
	if err != nil {
		t.Errorf("NewBosClient failed.")
	}
	c.Host = DebugHost

	file, err := os.Open(TestObjectName)
	if err != nil {
		t.Errorf("Open file failed")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	err = c.PutBucket(TestBukketName)
	_, err = c.PutObject(TestBukketName, TestObjectName, bytes.NewReader(content), "", "", nil)
	if err != nil {
		t.Errorf("PutObject failed.")
		t.Errorf(err.Error())
	}

	_, err = c.GetObject(TestBukketName, TestObjectName, 0, 0)
	if err != nil {
		t.Errorf("GetObject failed.")
		t.Errorf(err.Error())
	}

	c.DeleteObject(TestBukketName, TestObjectName)
}

func TestClean(t *testing.T) {
	os.Remove(TestObjectName)
	os.Remove(TestObjectName1)
}
