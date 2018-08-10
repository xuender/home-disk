import { Injectable, EventEmitter } from '@angular/core';
import { ModalController } from 'ionic-angular';
import { File } from '../../domain/file';
import { ImagePage } from '../../pages/image/image';
import { InfoPage } from '../../pages/info/info';

const PATH = 'assets/imgs/'
@Injectable()
export class PreviewProvider {

  srcDefault = `${PATH}unknown.png`
  itemChanged = new EventEmitter<any>();
  constructor(
    private modalCtrl: ModalController,
  ) {
    console.log('Hello PreviewProvider Provider');
  }
  preview(file: File) {
    this.modalCtrl
      .create(this.getPage(file), { file: file })
      .present();
  }
  private getPage(file: File): any {
    switch (file.type) {
      case 'image':
        return ImagePage
      default:
        return InfoPage
    }
  }
}
