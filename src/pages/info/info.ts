import { Component } from '@angular/core';
import { NavParams, ViewController } from 'ionic-angular';
import { File } from '../../domain/file';

@Component({
  selector: 'page-info',
  templateUrl: 'info.html',
})
export class InfoPage {

  file: File
  constructor(
    public navParams: NavParams,
    public viewCtrl: ViewController,
  ) {
    this.file = navParams.get('file')
  }
  cancel() {
    this.viewCtrl.dismiss()
  }
  ionViewDidLoad() {
    console.log('ionViewDidLoad InfoPage');
  }
  download() {
    const a = document.createElement("a")
    a.href = `/down/${this.file.id}`
    a.download = this.file.name
    a.click()
    a.remove()
  }
}
