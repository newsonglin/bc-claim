
function setGToken(val){
	if (typeof(Storage) !== "undefined") {
		/// Store
        localStorage.setItem("GTOKEN", val);
	} else {
		// Sorry! No Web Storage support..
	}
}



function getGToken(){
	// Retrieve
	return localStorage.getItem("GTOKEN");
}

