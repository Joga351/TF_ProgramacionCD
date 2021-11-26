<template>
    <div>
        <Header/>
            <h1>Lista de Criminales Registrados</h1>
            <div class="container izquierda">

                <button class="btn btn-primary" v-on:click="home()" >Home</button>
                <button class="btn btn-primary" v-on:click="listar()" style="margin-left: 10px" >Listar</button>
                <br><br>
                <button class="btn btn-dark margen" v-on:click="buscar()" >Buscar</button>
                <input type="text"  name="pagina" id="pagina" v-model="pagina">

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
                        <td>{{ persona.Cai }}</td>
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
            Personas2:null,
            pagina:'',
        }
    },
    components:{
        Header,
        Footer
    },
    methods:{
        home(){
            this.$router.push('/');
        },
            buscar(){
                let direccion2 = "http://localhost:3505/buscarDato?Index=1";
        axios.get(direccion2).then( data =>{this.Personas = data.data;
        });
            },

            listar(){
                let direccion = "http://localhost:3505/listarData";
        axios.get(direccion).then( data =>{this.Personas = data.data;
        });

            },
    }

}
</script>
<style  scoped>
    .izquierda{
        text-align: left;
    }
</style>