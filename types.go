package getdns

/*
 * getdns_return_t
 */

type ReturnType uint16

const (
	RETURN_GOOD                        ReturnType = 0
	RETURN_GENERIC_ERROR               ReturnType = 1
	RETURN_BAD_DOMAIN_NAME             ReturnType = 300
	RETURN_BAD_CONTEXT                 ReturnType = 301
	RETURN_CONTEXT_UPDATE_FAIL         ReturnType = 302
	RETURN_UNKNOWN_TRANSACTION         ReturnType = 303
	RETURN_NO_SUCH_LIST_ITEM           ReturnType = 304
	RETURN_NO_SUCH_DICT_NAME           ReturnType = 305
	RETURN_WRONG_TYPE_REQUESTED        ReturnType = 306
	RETURN_NO_SUCH_EXTENSION           ReturnType = 307
	RETURN_EXTENSION_MISFORMAT         ReturnType = 308
	RETURN_DNSSEC_WITH_STUB_DISALLOWED ReturnType = 309
	RETURN_MEMORY_ERROR                ReturnType = 310
	RETURN_INVALID_PARAMETER           ReturnType = 311
	RETURN_NOT_IMPLEMENTED             ReturnType = 312
)

const (
	DNSSEC_SECURE        = 400
	DNSSEC_BOGUS         = 401
	DNSSEC_INDETERMINATE = 402
	DNSSEC_INSECURE      = 403
	DNSSEC_NOT_PERFORMED = 404
)

const (
	NAMESPACE_DNS        = 500
	NAMESPACE_LOCALNAMES = 501
	NAMESPACE_NETBIOS    = 502
	NAMESPACE_MDNS       = 503
	NAMESPACE_NIS        = 504
)

type ResType uint16

const (
	RESOLUTION_STUB      ResType = 520
	RESOLUTION_RECURSING ResType = 521
)

const (
	GETDNS_REDIRECTS_FOLLOW        = 530
	GETDNS_REDIRECTS_DO_NOT_FOLLOW = 531
)

const (
	TRANSPORT_UDP = 1200
	TRANSPORT_TCP = 1201
	TRANSPORT_TLS = 1202
)

const (
	APPEND_NAME_ALWAYS                                    = 550
	APPEND_NAME_ONLY_TO_SINGLE_LABEL_AFTER_FAILURE        = 551
	APPEND_NAME_ONLY_TO_MULTIPLE_LABEL_NAME_AFTER_FAILURE = 552
	APPEND_NAME_NEVER                                     = 553
)

const (
	CONTEXT_CODE_NAMESPACES                    = 600
	CONTEXT_CODE_RESOLUTION_TYPE               = 601
	CONTEXT_CODE_FOLLOW_REDIRECTS              = 602
	CONTEXT_CODE_UPSTREAM_RECURSIVE_SERVERS    = 603
	CONTEXT_CODE_DNS_ROOT_SERVERS              = 604
	CONTEXT_CODE_DNS_TRANSPORT                 = 605
	CONTEXT_CODE_LIMIT_OUTSTANDING_QUERIES     = 606
	CONTEXT_CODE_APPEND_NAME                   = 607
	CONTEXT_CODE_SUFFIX                        = 608
	CONTEXT_CODE_DNSSEC_TRUST_ANCHORS          = 609
	CONTEXT_CODE_EDNS_MAXIMUM_UDP_PAYLOAD_SIZE = 610
	CONTEXT_CODE_EDNS_EXTENDED_RCODE           = 611
	CONTEXT_CODE_EDNS_VERSION                  = 612
	CONTEXT_CODE_EDNS_DO_BIT                   = 613
	CONTEXT_CODE_DNSSEC_ALLOWED_SKEW           = 614
	CONTEXT_CODE_MEMORY_FUNCTIONS              = 615
	CONTEXT_CODE_TIMEOUT                       = 616
	CONTEXT_CODE_IDLE_TIMEOUT                  = 617
)

const (
	RESPSTATUS_GOOD              = 900
	RESPSTATUS_NO_NAME           = 901
	RESPSTATUS_ALL_TIMEOUT       = 902
	RESPSTATUS_NO_SECURE_ANSWERS = 903
	RESPSTATUS_ALL_BOGUS_ANSWERS = 904
)

const (
	BAD_DNS_CNAME_IN_TARGET               = 1100
	BAD_DNS_ALL_NUMERIC_LABEL             = 1101
	BAD_DNS_CNAME_RETURNED_FOR_OTHER_TYPE = 1102
)
