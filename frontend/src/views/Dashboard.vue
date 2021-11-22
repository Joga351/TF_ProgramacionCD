<template>
    <div>
        <Header/>

            <div class="container izquierda">

                <button class="btn btn-primary" v-on:click="Buscar()" >Buscar</button>
                <br><br>


                <table class="table table-hover">
                <thead>
                    <tr>
                        <th scope="col">Nº ORDEN</th>
                        <th scope="col">DISTRITO</th>
                        <th scope="col">EDAD</th>
                        <th scope="col">TRABAJO</th>
                        <th scope="col">VINCULO</th>
                        <th scope="col">TIPO DE VIOLENCIA</th>
                        <th scope="col">CONSUMO DE ALCOHOL</th>
                        <th scope="col">CONSUMO DE NICOTINA</th>
                        <th scope="col">CONSUMO DE DROGA</th>
                        <th scope="col">ADICCION</th>
                        <th scope="col">RIESGO</th>
                        <th scope="col">MES</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="persona in Personas" :key="persona.index">
                        <th scope="row">{{ persona.index }}</th>
                        <td>{{ persona.CAI }}</td>
                        <td>{{ persona.Edad }}</td>
                        <td>{{ persona.Trabajo }}</td>
                        <td>{{ persona.Vinculo }}</td>
                        <td>{{ persona.TipoVio }}</td>
                        <td>{{ persona.ConsumeAlcohol }}</td>
                        <td>{{ persona.Fuma }}</td>
                        <td>{{ persona.ConsumeDroga }}</td>
                        <td>{{ persona.Adiccion }}</td>
                        <td>{{ persona.RiesgoPresuntivo }}</td>
                        <td>{{ persona.Mes }}</td>
                    </tr>
            
                </tbody>
                </table>

            </div>

        <Footer />
    </div>
</template>
<script>
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
import axios from 'axios';
export default {
    name:"Dashboard",
    data(){
        return {
            Personas:null,
        }
    },
    components:{
        Header,
        Footer
    },
    
    mounted:function(){
        let direccion = "http://localhost:3505/listarData";
        axios.get(direccion).then( data =>{this.Personas = data.data;
        });
    }
}
</script>
<style  scoped>
    .izquierda{
        text-align: left;
    }
</style>