import Vue from 'vue';
import Component from 'vue-class-component';
import httpService from '../../../services/http.service';
import navService from '../../../services/nav.service';
import userService from '@/services/user.service';

@Component
export default class Login extends Vue {
  public validForm:boolean = false;

  public pseudo:string = '';

  public password:string = '';

  public showPassword:boolean = false;

  public pseudoRules:any = [
    (v:string) => !!v || 'Pseudo is required',
  ];

  public passwordRules:any = [
    (v:string) => !!v || 'Password is required',
  ];

  public displayAlert:boolean = false;

  constructor() {
    super();
  }

  public goRegister() {
    navService.changeView('register');
  }

  public mounted() {
    this.justPing();
  }

  public isValid():boolean {
    return this.validForm && !this.displayAlert;
  }

  public connection() {
    let headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'username', this.pseudo);
    headers = httpService.appendHeaders(headers, 'password', btoa(`${this.pseudo}:${this.password}`));
    httpService.post('login', {}, headers, (response:any) => {
      userService.connectionDone(response.data.token, response.data.user);
    	navService.changeView('/');
    }, (error:any) => {
      this.displayAlert = true;
      console.log(error);
    });
  }

  public getURL() { return httpService.getDisplayServerURL(); }

  public justPing() {
    this.displayAlert = false;
    httpService.get('ping', httpService.getDefaultHeaders(), (response:any) => {
      console.log(response.data);
    }, (error:any) => {
      this.displayAlert = true;
    });
  }
}
