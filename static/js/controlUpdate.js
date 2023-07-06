
  async function actualizarControl() {
   let datos = {};
     datos.descripcion = document.getElementById('txtDescripcion').value;
      const url = new URL(document.URL);
      const searchParams = url.searchParams;

      let id = searchParams.get('control_id');

          const request = await fetch(`control/${id}`, {
              method: 'PUT',
              body: JSON.stringify(datos),
              headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json'
              },
          })
        if (request.status == 200) {
            alert("Control actualizado correctamente");
            window.location.href = 'controles.html';
        }
  }
