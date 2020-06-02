format:  
	bash ./scripts/format.sh

check: format
	bash ./scripts/check.sh
testing: check  
	bash ./scripts/test.sh 

 	
			 
