
async function updatePiso(){
    let datos = {};

    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    let pisoid = searchParams.get('piso_id');
    
    
    numero = document.getElementById('NroPiso').value;
    datos.numero = Number(numero);

    const request = await fetch("http://3.83.152.157:5000/pisos/" + pisoid, {
               method: 'PUT',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("piso actualizado correctamente");
         }else{
             alert("Error actualizando el piso");
         }
         
         window.location.href = `pisos.html`;
}