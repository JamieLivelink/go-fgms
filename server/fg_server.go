
package server


const SUCCESS                 = 0
const ERROR_COMMANDLINE       = 1
const ERROR_CREATE_SOCKET     = 2
const ERROR_COULDNT_BIND      = 3
const ERROR_NOT_LISTENING     = 4
const ERROR_COULDNT_LISTEN    = 5

// other constants
const MAX_PACKET_SIZE         = 1024
const UPDATE_INACTIVE_PERIOD  = 1
const MAX_TELNETS             = 5
const RELAY_MAGIC             = 0x53464746    // GSGF




		
// http://gitorious.org/fgms/fgms-0-x/blobs/master/src/server/fg_server.cxx#line167
type FG_SERVER struct {

  /*typedef union
  {
    uint32_t    complete;
    int16_t     High;
    int16_t     Low;
  } converter; 
  converter*    tmp; */
  Initialized bool
  ReinitData bool
  ReinitTelnet bool
  Listening bool
  
  ServerName  string
  BindAddress string
  ListenPort int
  
  PlayerExpires int
  
  IamHUB bool
 
  //Loglevel            = SG_INFO;
  DataSocket int
  TelnetPort int
  NumMaxClients int
  PlayerIsOutOfReach int // nautical miles
  NumCurrentClients int
  IsParent bool
  MaxClientID int

  //tmp                   = (converter*) (& PROTO_VER);
  //ProtoMinorVersion   = tmp->High;
  //ProtoMajorVersion   = tmp->Low;
  //LogFileName         = DEF_SERVER_LOG; // "fg_server.log";
  //wp                  = fopen("wp.txt", "w");
  //BlackList           = map<uint32_t, bool>();
  //RelayMap            = map<uint32_t, string>();
  IsTracked bool
  Tracker int
  //UpdateSecs          = DEF_UPDATE_SECS;
  // clear stats - should show what type of packet was received
  PacketsReceived int
  TelnetReceived int 
  BlackRejected int
  PacketsInvalid int //     = 0;  // invalid packet
  UnknownRelay int //       = 0;  // unknown relay
  RelayMagic  int //        = 0;  // relay magic packet
  PositionData int //         = 0;  // position data packet
  NotPosData int    //     = 0;
  // clear totals
  MT_PacketsReceived int
  MT_BlackRejected int
  MT_PacketsInvalid int
  MT_UnknownRelay int
  MT_PositionData int
  MT_TelnetReceived int
  MT_RelayMagic int
  MT_NotPosData int
  
  CrossFeedFailed int
  CrossFeedSent int
  
  MT_CrossFeedFailed int
  MT_CrossFeedSent int
  TrackerConnect int
  TrackerDisconnect int
  TrackerPostion int // Tracker messages queued
  //pthread_mutex_init( &m_PlayerMutex, 0 );
} 

// Consruct and return pointer to new FG_SERVER instance
func NewFG_SERVER() *FG_SERVER {
	ob := new(FG_SERVER)
	// set other defaults here
	return ob
}

func (me *FG_SERVER) SetServerName(name string){
	me.ServerName = name
}
func (me *FG_SERVER) SetBindAddress(addr string){
	me.BindAddress = addr
}

func (me *FG_SERVER) SetDataPort(port int){
	me.ListenPort = port
	me.ReinitData = true
}

func (me *FG_SERVER) SetTelnetPort(port int){
	me.TelnetPort = port
	me.ReinitTelnet = true
}



// Set nautical miles two players must be apart to be out of reach
func (me *FG_SERVER) SetOutOfReach(nm int){
	me.PlayerIsOutOfReach = nm
}

//  Set time in seconds. if no packet arrives from a client
//  within this time, the connection is dropped.  
func (me *FG_SERVER) SetPlayerExpires(secs int){
	me.PlayerExpires = secs
}

// Set if we are running as a Hubserver
func (me *FG_SERVER) SetHub(am_hub bool){
	me.IamHUB = am_hub
}






//////////////////////////////////////////////////
// mT_Relay - Type of list of relays
type MT_Relay struct {
	Name string
	Address string // TODO = netAddress  Address
}

