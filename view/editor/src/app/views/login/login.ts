import Vue from 'vue';
import Component from 'vue-class-component';
import httpService from '../../../services/http.service';

@Component
export default class Login extends Vue {

  public validForm:boolean = false;
  public pseudo:string = '';
  public password:string = '';
  public showPassword:boolean = false;
  public pseudoRules:any = [
    (v:string) => !!v || 'Pseudo is required'
  ];
  public passwordRules:any = [
    (v:string) => !!v || 'Password is required'
  ];

  public displayAlert:boolean = false;

  constructor(){
    super();
  }

  public mounted(){
    this.justPing();
  }

  public isValid():boolean{
    return this.validForm;
  }

  public connection(){
    alert('TODO');
  }

  public getURL(){ return httpService.getDisplayServerURL(); }

  public justPing(){
    this.displayAlert = false;
    httpService.get('ping', (response:any) => {
      console.log(response.data);
    }, (error:any) => {
      this.displayAlert = true;
    });
  }

}
