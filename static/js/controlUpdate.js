async function updateControl(){
    let datos = {};

    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    let control_id = searchParams.get('control_id');
    
    
    datos.descripcion = document.getElementById('txtDescripcion').value;
    datos.tipo = document.getElementById('selectTipo').value;
    

    const request = await fetch("http://34.192.187.56:5000/controles/" + control_id, {
               method: 'PUT',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("controles actualizado correctamente");
         }else{
             alert("Error actualizando el control");
         }
         
         window.location.href = `controles.html`;
}