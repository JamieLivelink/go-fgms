package fgms

import (
	//"time"
)

type FG_Player struct {
	//public:
	Origin string
	Address *NetAddress    
	Callsign string
	Passwd string
	ModelName string
	JoinTime int64 //time.Time
	Timestamp int64 //time.Time
	LastPos *Point3D
	LastOrientation *Point3D
	IsLocal bool
	Error string //;    // in case of errors
	HasErrors bool
	ClientID int
	LastRelayedToInactive int64 //time.Time
	
	// Packets recieved from client 
	PktsReceivedFrom uint  
	
	// Packets sent to client 
	PktsSentTo uint       
	
	//  Packets from client sent to other players/relays 
	PktsForwarded uint    

	//FG_Player ();
	//FG_Player ( const FG_Player& P);
	//~FG_Player ();
	//void operator =  ( const FG_Player& P );
	//private:
	//void assign ( const FG_Player& P );
} // FG_Player

func NewFG_Player() *FG_Player {
	ob := new(FG_Player)
	return ob
}

