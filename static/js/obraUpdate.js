async function updateObra(){
    let datos = {};

    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    let obraid = searchParams.get('obra_id');
    
    
    datos.nombre = document.getElementById('txtNombre').value;
    

    const request = await fetch("http://localhost:5000/obras/" + obraid, {
               method: 'PUT',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("obra actualizada correctamente");
         }else{
             alert("Error actualizando la obra");
         }
         
         window.location.href = `obras.html`;
}