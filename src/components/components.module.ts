import { NgModule } from '@angular/core';
import { IonicModule } from 'ionic-angular'
import { FileUploadModule } from 'ng2-file-upload'
//import { NgxUploaderModule } from 'ngx-uploader'

import { HdUploadComponent } from './hd-upload/hd-upload';
@NgModule({
  declarations: [HdUploadComponent],
  imports: [
    IonicModule,
    FileUploadModule,
  ],
  exports: [
    HdUploadComponent,
    FileUploadModule,
  ]
})
export class ComponentsModule { }
