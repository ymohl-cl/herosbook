
import router from '@/router';

/**
 * Service to manage navigation between views and rewrite url
 */
class CNavService {

  public changeView(route:string):void{
    router.push(route);
  }

  public replaceView(route:string):void{
    router.replace(route);
  }

  public isSameRoute(route:string):boolean{
    return router.currentRoute.name == route;
  }
}

const NavService = new CNavService();
export default NavService; 