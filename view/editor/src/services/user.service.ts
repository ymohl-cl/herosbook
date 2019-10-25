import HttpService from './http.service';

/**
 * Service to manage connection and user
 */
class CUserService {
  private token:string;

  private user:any;

  constructor() {
    this.token = '';
  }

  public connectionDone(token:string, user:any) {
    this.token = token;
    this.user = user;
  }

  public isConnected():boolean {
    return this.token !== '';
  }

  public getPseudo():string {
    if (this.isConnected() && this.user != null
		&& this.user.pseudo != null) {
      return this.user.pseudo;
    }
    return '';
  }

  public getToken():string {
    return this.token;
  }
}

const UserService = new CUserService();
export default UserService;
