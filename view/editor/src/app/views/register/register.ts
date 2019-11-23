import Vue from 'vue';
import Component from 'vue-class-component';
import httpService from '@/services/http.service';
import navService from '@/services/nav.service';

@Component
export default class Register extends Vue {
  public validForm:boolean = false;

  public genres:string[] = [
    'Male',
    'Female',
    'Anothers',
  ];

  public pseudo:string = '';

  public password:string = '';

  public lastName:string = '';

  public firstName:string = '';

  public email:string = '';

  public age:number = 0;

  public genre:string = '';

  public showPassword:boolean = false;

  public pseudoRules:any = [
    (v:string) => !!v || 'Pseudo is required',
  ];

  public passwordRules:any = [
    (v:string) => !!v || 'Password is required',
  ];

  public emailRules:any = [
    (v:string) => !!v || 'Email is required',
    (v:string) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
  ];

  public lastNameRules:any = [
    (v:string) => !!v || 'Last name is required',
  ];

  public firstNameRules:any = [
    (v:string) => !!v || 'First name is required',
  ];

  public ageRules:any = [
    (v:number) => v > 0 || 'Age not null or negative',
  ];

  public genreRules:any = [
    (v:string) => !!v || 'Genre is required',
  ];

  public displayAlert:boolean = false;

  public mounted() {
    this.justPing();
  }

  public goLogin() {
    navService.changeView('login');
  }

  public isValid():boolean {
    return this.validForm && !this.displayAlert;
  }

  public register() {
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'password', btoa(`${this.pseudo}:${this.password}`));
    httpService.post('register', {
      pseudo: this.pseudo,
      name: this.firstName,
      lastName: this.lastName,
      age: this.age,
      genre: this.genre,
      email: this.email,
    }, headers, (response:any) => {
      console.log(response.data);
    }, (error:any) => {
      this.displayAlert = true;
      console.log(error);
    });
  }

  public getURL():string {
    return httpService.getDisplayServerURL();
  }

  public justPing() {
    this.displayAlert = false;
    httpService.get('ping', httpService.getDefaultHeaders(), (response:any) => {
      console.log(response.data);
    }, (error:any) => {
      this.displayAlert = true;
    });
  }
}
