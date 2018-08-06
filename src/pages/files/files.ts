import { Component } from '@angular/core';
import { NavController, NavParams, ModalController } from 'ionic-angular';
import { File } from '../../domain/file'
import { FilesProvider } from '../../providers/files/files'
import { ImagePage } from '../image/image'

@Component({
  selector: 'page-files',
  templateUrl: 'files.html',
})
export class FilesPage {
  constructor(
    public navCtrl: NavController,
    public navParams: NavParams,
    public modalCtrl: ModalController,
    public files: FilesProvider,
  ) {
  }
  doInfinite(infiniteScroll) {
    this.files.loadDay()
      .then(r => infiniteScroll.complete())
  }
  onSelectFile(file: File) {
    this.modalCtrl.create(this.getPage(file.type), { file: file })
      .present();
  }
  private getPage(type: string): any {
    switch (type) {
      case 'image':
        return ImagePage
      default:
        return ImagePage
    }
  }
}
