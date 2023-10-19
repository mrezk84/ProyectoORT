async function updateFormulario(){
    let datos = {};

    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    let formulario_id = searchParams.get('formulario_id');
    
    
    datos.nombre = document.getElementById('txtNombre').value;
    datos.informacion = document.getElementById('txtInformacion').value;
    

    const request = await fetch("http://3.83.152.157:5000/formularios/" + formulario_id, {
               method: 'PUT',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("formulario actualizado correctamente");
         }else{
             alert("Error actualizando el formulario");
         }
         
         window.location.href = `formularios.html`;
}