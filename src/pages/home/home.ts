import { Component } from '@angular/core';
import { Platform, NavController } from 'ionic-angular';
import { FilesProvider } from '../../providers/files/files';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {
  public pc = true
  constructor(
    public navCtrl: NavController,
    public plt: Platform,
    private filesProvider: FilesProvider
  ) {
    this.pc = !this.plt.is('mobile') && !this.plt.is('mobileweb')
  }
  onUpload(ok: boolean) {
    setTimeout(() => this.filesProvider.reset(), 1000)
  }
}
