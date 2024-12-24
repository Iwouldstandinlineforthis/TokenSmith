package classes

var RedirURI string
var ClientID string       //= "1fec8e78-bce4-4aaf-ab1b-5451cc387264"
var ResourceURL string    // = "https://graph.microsoft.com/.default"
var RefResourceURL string // = "https://graph.microsoft.com/"
var Scope string          // = "openid offline_access"
var UserAgent string      // = "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0"
var RefreshToken string   // "0.A.."

// var DefaultUserAgent string = "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0"

const AuthorizeV2Endpoint string = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
const TokenV2Endpoint string = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
const TokenV1Endpoint string = "https://login.microsoftonline.com/common/oauth2/token"

// const TokenV2Endpoint string = "http://localhost:8000/post"
