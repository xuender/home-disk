import { Component } from '@angular/core';
import { Platform, NavController } from 'ionic-angular';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {
  public pc = true
  constructor(
    public navCtrl: NavController,
    public plt: Platform
  ) {
    this.pc = !this.plt.is('mobile') && !this.plt.is('mobileweb')
  }
}
