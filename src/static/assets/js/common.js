/*
@author https://github.com/macek/jquery-serialize-object
*/
$.fn.serializeObject = function() { 
    var obj = null; 
    try { 
        if(this[0].tagName && this[0].tagName.toUpperCase() == "FORM" ) {
             var arr = this.serializeArray(); 
             if(arr){ 
                 obj = {}; 
                 $.each(arr, function() { 
                     obj[this.name] = this.value; 
                    }); 
                }
            } 
        }catch(e) { 
            alert(e.message); 
        }finally {

        } 
        return obj; 
    }

