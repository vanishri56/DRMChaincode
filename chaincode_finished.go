package main

  import(
    "encoding/json";
    "errors";
    "fmt";
	"strconv";
   
"github.com/hyperledger/fabric/core/chaincode/shim";
   "github.com/hyperledger/fabric/core/crypto/primitives"
    )
    
/* DRMChaincode is a high level smart contract for usage of business partners to manage the digital copyrights */
    
type DRMChaincode struct {
}

func main() {
	err := shim.Start(new(DRMChaincode))
	if err != nil {
		fmt.Printf("Error starting DRM chaincode: %s", err)
	}
}

/* ConsumerDetails is for storing User Details*/

/*ownership could be Purchase or Lease*/
/*period of lease - as in number of days*/


type DRMConsumerDetails struct{	
	ConsumerId string;
	FirstName string;
	LastName string;
	Dob string; 
 	RetailerSubid string;
  	Typeofownership string;
	Periodoflease string;
	PaymentId string 
	}

type AddressDetails struct{
	AddressId string; 
	Email string; 
	Address1 string;
	Address2 string; 
	City string; 
	State string;
	Zip string;
	Country string; 	}
	
  type DRM_ProducerDetails struct {
    ProducerId string;
    ProducerName string;
    ProductType string;
    ProductName string; 
    RegistrationId string;
    ProductId string;
    DRMAuthorityName string;
    }

  type DRM_RetailerDetails struct {
    ProducerId string ;
    ProductId string ;  
    ProductType string ;    
ProductName string ; 
    RetailerSubid string ;
    }
  
  
  type Transaction struct{	
	TransactionId string;
	TimeStamp string ;
	DRMId string ;
	ConsumerId string ;
	Transactiontype string ;
  	ProducerId string ;	
	ProductId string ;
	RetailerSubid ;
}

  
// Init initializes the smart contracts
func (t *DRMchaincode) init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}


	// Check if table already exists
	err := stub.GetTable("DRMConsumerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, err
	}


	// Create application Table DRMConsumerDetails

	err = stub.CreateTable("DRMConsumerDetails",[]*shim.ColumnDefinition{
&shim.ColumnDefinition{Name: "ConsumerId", Type:shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "FirstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Lastname", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "RetailerSubid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Typeofownership", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentId", Type: shim.ColumnDefinition_STRING, Key: true},
}) 

	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
		}
      
      
      // Check if table already exists
	err = stub.GetTable("AddressDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, err
	}
	
	// Create AddressDetails Table
	err = stub.CreateTable("AddressDetails",[]*shim.ColumnDefinition{
&shim.ColumnDefinition{Name: "AddressId", Type:shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Address1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Address2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "state", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}


      // Check if table already exists 
	err = stub.GetTable("DRM_ProducerDetails")
	if err == nil {
		return nil, errors.New("Table already exists.")
	}
	
		// Create application Table
	err = stub.CreateTable(“DRM_ProducerDetails”, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ProducerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ProducerName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "RegistrationId", Type: shim.ColumnDefinition_STRING, Key: True},
		&shim.ColumnDefinition{Name: "ProductId", Type: shim.ColumnDefinition_STRING, Key: Trye},
		&shim.ColumnDefinition{Name: "DRMAuthorityName", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}


      // Check if table already exists
	err = stub.GetTable("DRM_RetailerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, err
	}
	
		// Create application Table
	err = stub.CreateTable("DRM_RetailerDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ProducerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ProductId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ProductType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProductName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "RetailerSubid", Type :shim.ColumnDefinition_STRING, Key: true}
		
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

	
		// Create application Table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "TransactionId", Type: shim.ColumnDefinition_STRING, Key: true}, &shim.ColumnDefinition{Name: "TimeStamp", Type: shim.ColumnDefinition_STRING, Key: false},&shim.ColumnDefinition{Name: "DRMId", Type: shim.ColumnDefinition_STRING, Key: True}, &shim.ColumnDefinition{Name: "ConsumerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Transactiontype", Type :shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ProducerId", Type :shim.ColumnDefinition_STRING, Key: true}, &shim.ColumnDefinition{Name:"ProductId", 
Type :shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "RetailerSubid", Type :shim.ColumnDefinition_STRING, Key: true},
		
		
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	// setting up the users role
	stub.PutState("user_type1_1", []byte("Producer"))
	stub.PutState("user_type1_2", []byte("Retailer"))
	stub.PutState("user_type1_3", []byte("DRMAuthority"))
	stub.PutState("user_type1_4", []byte("Consumer"))	
	
	return nil, err
}

//registerUser to Producer registers producet with the DRM
func (t *DRMChaincode) register(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


if len(args) != 9 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 9. Got: %d.", len(args))
}
		
	ProducerId:args[1]
	ProducerName:args[2]
	ProductType:args[3]
	ProductName:args[4]
	RetailerSubid:args[4]
	RegistrationId:args[5]
  	ProductId:args[6]
  	DRMAuthorityName:args[7]
  
assignerOrg1, err := stub.GetState(args[8])
assignerOrg := string(assignerOrg1)

createdBy:=assignerOrg
totalPoint:="0"

	// Insert a row
err := stub.InsertRow(DRM_ProducerDetails, shim.Row{Columns: []*shim.Column{
				{Value: &shim.Column_String_{String_:ProducerId}},
		{Value: &shim.Column_String_{String_: ProducerName}},
		{Value: &shim.Column_String_{String_: ProductType}},
	{Value: &shim.Column_String_{String_: RetailerSubid}},
	{Value: &shim.Column_String_{String_: RegistrationId}},
	{Value: &shim.Column_String_{String_: ProductId}},
	{Value: &shim.Column_String_{String_: DRMAuthorityName}},
				
				
			},
})


		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, err


}
	// Get all rows pertaining to this DRM_ProducerDetails
	Rows, err := GetRow{DRM_ProducerDetails,[]shim,.column{}) }}



	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, err
	}
// Invoke invokes the chaincode
func (t *DRMChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {


	if function == "register" {
		t := DRMChaincode{}
		return t.register(stub, args)	
	} 
	}


	return nil, errors.New("Invalid invoke function name.")


	}


/* ============================================================================================================================
	// Invoke - Our entry point for Invocations
	// ===========================================================================================================================

func (t *DRMChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
		fmt.Println("invoke is running " + function)

		// Handle different functions
		if function == "Init" {               //initialize the chaincode
			return t.Init(stub, "init", args)
		} else if function == "delete" {   //deletes an entity from its state
			return t.Delete(stub, args)
		} else if function == "write" {       //writes a value to the chaincode state
			return t.Write(stub, args)
		} else if function == "init" { //create a new marble
			return t.init(stub, args)
		}
		fmt.Println("invoke did not find func: " + function) //error

		return nil, errors.New("Received unknown function invocation")
	}

*/













