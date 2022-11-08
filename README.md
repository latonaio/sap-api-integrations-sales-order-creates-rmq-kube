# sap-api-integrations-sales-order-creates-rmq-kube  
sap-api-integrations-sales-order-creates-rmq-kube は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で受注データを登録するマイクロサービスです。  
sap-api-integrations-sales-order-creates-rmq-kube には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-sales-order-creates-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview  

## 動作環境  
sap-api-integrations-sales-order-creates-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-sales-order-creates-rmq-kube は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-order-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview  
* APIサービス名(=baseURL): API_SALES_ORDER_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-order-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_SalesOrder（受注）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"HeaderItem" が指定されています。

```
"api_schema": "SAPSalesOrdercreates-rmq-kube",
"accepter": ["HeaderItem"],
"sales_order": "",
"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPSalesOrdercreates-rmq-kube",
"accepter": ["All"],
"sales_order": "",
"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostSalesOrder(
	header             *requests.Header,
	item               *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
				func() {
					c.Item(item)
					wg.Done()
				}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 受注 の ヘッダデータ/明細データ が登録された結果の JSON の例です。  
以下の項目のうち、"SalesOrder" ～ "BillingDocumentDate" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-order-creates-rmq-kube/SAP_API_Caller/caller.go#L59",
	"function": "sap-api-integrations-sales-order-creates-rmq-kube/SAP_API_Caller.(*SAPAPICaller).HeaderItem",
	"level": "INFO",
	"message": {
		"SalesOrder": "14",
		"SalesOrderType": "OR1",
		"SalesOrganization": "0001",
		"DistributionChannel": "01",
		"OrganizationDivision": "01",
		"SalesGroup": "",
		"SalesOffice": "",
		"SalesDistrict": "",
		"SoldToParty": "1",
		"CreationDate": null,
		"LastChangeDate": null,
		"ExternalDocumentID": "",
		"LastChangeDateTime": null,
		"PurchaseOrderByCustomer": "Test",
		"CustomerPurchaseOrderDate": "2022-09-15",
		"SalesOrderDate": null,
		"TotalNetAmount": null,
		"TransactionCurrency": "",
		"SDDocumentReason": "",
		"PricingDate": null,
		"RequestedDeliveryDate": null,
		"ShippingCondition": "",
		"CompleteDeliveryIsDefined": false,
		"ShippingType": "",
		"HeaderBillingBlockReason": "",
		"DeliveryBlockReason": "",
		"IncotermsClassification": "",
		"CustomerPriceGroup": "",
		"PriceListType": "",
		"CustomerPaymentTerms": "",
		"PaymentMethod": "",
		"ReferenceSDDocument": "",
		"ReferenceSDDocumentCategory": "",
		"CustomerAccountAssignmentGroup": "",
		"AccountingExchangeRate": null,
		"CustomerGroup": "",
		"AdditionalCustomerGroup1": "",
		"AdditionalCustomerGroup2": "",
		"AdditionalCustomerGroup3": "",
		"AdditionalCustomerGroup4": "",
		"AdditionalCustomerGroup5": "",
		"CustomerTaxClassification1": "",
		"TotalCreditCheckStatus": "",
		"to_Item": {
			"results": [
				{
					"SalesOrder": "",
					"SalesOrderItem": "10",
					"SalesOrderItemCategory": null,
					"SalesOrderItemText": null,
					"PurchaseOrderByCustomer": null,
					"Material": "21",
					"MaterialByCustomer": null,
					"PricingDate": null,
					"RequestedQuantity": "1",
					"RequestedQuantityUnit": null,
					"ItemGrossWeight": "1",
					"ItemNetWeight": "1",
					"ItemWeightUnit": null,
					"ItemVolume": "1",
					"ItemVolumeUnit": null,
					"TransactionCurrency": null,
					"NetAmount": null,
					"MaterialGroup": null,
					"MaterialPricingGroup": null,
					"Batch": null,
					"ProductionPlant": null,
					"StorageLocation": null,
					"DeliveryGroup": null,
					"ShippingPoint": null,
					"ShippingType": null,
					"DeliveryPriority": null,
					"IncotermsClassification": "",
					"ProductTaxClassification1": null,
					"MatlAccountAssignmentGroup": null,
					"CustomerPaymentTerms": null,
					"CustomerGroup": null,
					"SalesDocumentRjcnReason": null,
					"ItemBillingBlockReason": null,
					"WBSElement": null,
					"ProfitCenter": null,
					"AccountingExchangeRate": null,
					"ReferenceSDDocument": null,
					"ReferenceSDDocumentItem": null,
					"SDProcessStatus": null,
					"DeliveryStatus": null,
					"OrderRelatedBillingStatus": null
				},
				{
					"SalesOrder": "",
					"SalesOrderItem": "20",
					"SalesOrderItemCategory": null,
					"SalesOrderItemText": null,
					"PurchaseOrderByCustomer": null,
					"Material": "21",
					"MaterialByCustomer": null,
					"PricingDate": null,
					"RequestedQuantity": "1",
					"RequestedQuantityUnit": null,
					"ItemGrossWeight": "1",
					"ItemNetWeight": "1",
					"ItemWeightUnit": null,
					"ItemVolume": "1",
					"ItemVolumeUnit": null,
					"TransactionCurrency": null,
					"NetAmount": null,
					"MaterialGroup": null,
					"MaterialPricingGroup": null,
					"Batch": null,
					"ProductionPlant": null,
					"StorageLocation": null,
					"DeliveryGroup": null,
					"ShippingPoint": null,
					"ShippingType": null,
					"DeliveryPriority": null,
					"IncotermsClassification": "",
					"ProductTaxClassification1": null,
					"MatlAccountAssignmentGroup": null,
					"CustomerPaymentTerms": null,
					"CustomerGroup": null,
					"SalesDocumentRjcnReason": null,
					"ItemBillingBlockReason": null,
					"WBSElement": null,
					"ProfitCenter": null,
					"AccountingExchangeRate": null,
					"ReferenceSDDocument": null,
					"ReferenceSDDocumentItem": null,
					"SDProcessStatus": null,
					"DeliveryStatus": null,
					"OrderRelatedBillingStatus": null
				}
			]
		}
	},
	"time": "2022-09-23T19:14:35+09:00"
}
```

