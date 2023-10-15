/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2023. All rights reserved.
 * Licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Description: declaration of logger tool interfaces.
 */
#ifndef CERT_CONFIG_H
#define CERT_CONFIG_H

#define CERT_MANAGER_DEPLOY_PATH "/usr/bin/certmanager"
#define CERT_MANAGER_DEPLOY_USER "root"

/*
 * defines the public key for verifying the imported certification.
 */
const char g_root_public_key[] = {
/* add public_key len 550*/
    0x30, 0x82, 0x02, 0x22, 0x30, 0x0D, 0x06, 0x09, 0x2A, 0x86, 
    0x48, 0x86, 0xF7, 0x0D, 0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 
    0x82, 0x02, 0x0F, 0x00, 0x30, 0x82, 0x02, 0x0A, 0x02, 0x82, 
    0x02, 0x01, 0x00, 0xBF, 0x96, 0x1D, 0x26, 0xE4, 0x7A, 0xD0, 
    0x9A, 0xF1, 0x49, 0xC7, 0xAA, 0x24, 0xA8, 0x0E, 0xDE, 0x7F, 
    0x24, 0xDC, 0x53, 0xFE, 0x9E, 0xD8, 0x0F, 0x6A, 0x3F, 0x82, 
    0x58, 0xB2, 0x24, 0x5E, 0xA5, 0x16, 0xD8, 0x4B, 0x6B, 0x65, 
    0x23, 0xC3, 0x6B, 0x44, 0x8B, 0x43, 0x2B, 0x40, 0x40, 0x5E, 
    0x25, 0xCF, 0x2A, 0x6A, 0x12, 0x44, 0xBE, 0x99, 0x8C, 0xD2, 
    0xDC, 0xA4, 0x1E, 0xB8, 0xBE, 0x87, 0xC3, 0xB8, 0x61, 0x61, 
    0x7D, 0xA8, 0x22, 0x78, 0xD9, 0x9B, 0x28, 0x73, 0x90, 0xB7, 
    0x42, 0x58, 0x33, 0x48, 0x37, 0xC7, 0xA4, 0x26, 0x74, 0xE8, 
    0xBA, 0x0C, 0x5E, 0xD6, 0x13, 0xCF, 0x05, 0x1D, 0x6C, 0x9D, 
    0xC3, 0x65, 0x96, 0x0D, 0xD9, 0x00, 0x40, 0x14, 0x5D, 0xBE, 
    0x8D, 0x05, 0x24, 0xA8, 0x46, 0xDD, 0xE5, 0x4B, 0x6A, 0x84, 
    0xB2, 0x25, 0xF7, 0xBA, 0x34, 0x0D, 0x57, 0x7C, 0x7A, 0x1C, 
    0xA4, 0x5E, 0x6E, 0x4C, 0x8E, 0xB0, 0xA2, 0x43, 0x5F, 0x52, 
    0x24, 0x41, 0x62, 0x0C, 0xD9, 0x86, 0xF9, 0xF1, 0xFD, 0xDF, 
    0x9D, 0x47, 0x0C, 0x0A, 0xE0, 0x74, 0x99, 0xDC, 0xBB, 0x04, 
    0xA8, 0x18, 0x95, 0xA6, 0xFF, 0xAA, 0xAE, 0xBB, 0xA4, 0x54, 
    0x79, 0x7C, 0xD6, 0x2E, 0x81, 0xB6, 0xED, 0x37, 0x6B, 0xD0, 
    0xFE, 0x0C, 0x70, 0xB7, 0xC7, 0xE0, 0x23, 0x9D, 0x9F, 0x7C, 
    0x34, 0x14, 0x11, 0x68, 0xCE, 0x27, 0x6E, 0xB1, 0xDD, 0x96, 
    0x5E, 0x55, 0xB1, 0xDB, 0xBD, 0xEC, 0x12, 0x23, 0xAA, 0xF0, 
    0x1C, 0x6D, 0xA0, 0xB5, 0x76, 0x3B, 0x8F, 0x7C, 0x50, 0x98, 
    0x30, 0xCD, 0x2E, 0x0E, 0x19, 0xD9, 0x69, 0xE7, 0xE1, 0x0B, 
    0xC3, 0x98, 0xE2, 0x89, 0x8C, 0xE3, 0xF4, 0xE7, 0x5D, 0x34, 
    0x12, 0x82, 0x9E, 0xA5, 0x11, 0x8E, 0xCB, 0x34, 0xB4, 0x48, 
    0xE4, 0x4A, 0xB7, 0x37, 0xCB, 0xC1, 0xC9, 0xDC, 0xF0, 0xEC, 
    0x12, 0xA9, 0x02, 0x90, 0x03, 0xF8, 0xE5, 0x21, 0xA4, 0x7E, 
    0x0A, 0x27, 0x22, 0x98, 0x03, 0x66, 0xAE, 0xCA, 0xB2, 0x2F, 
    0x7F, 0xC6, 0x0E, 0x84, 0x5F, 0x59, 0xA3, 0x66, 0x9D, 0x01, 
    0x52, 0x08, 0xA4, 0xA5, 0xA2, 0xCB, 0x84, 0xE0, 0x8C, 0x3F, 
    0x30, 0xFB, 0x1C, 0x73, 0xA2, 0x6E, 0x72, 0x6A, 0xD5, 0xC1, 
    0x9E, 0xE2, 0x0D, 0x0B, 0x98, 0xE1, 0xAA, 0x2A, 0xFA, 0x97, 
    0x75, 0xE0, 0xB8, 0x86, 0x47, 0x53, 0xAF, 0xDC, 0xC4, 0xA6, 
    0xA9, 0xE2, 0xEB, 0x7D, 0xAE, 0x88, 0x60, 0xD9, 0x51, 0xE9, 
    0x12, 0xCA, 0x4F, 0x50, 0xAC, 0x00, 0xF0, 0x33, 0x78, 0x59, 
    0xCE, 0xD8, 0x1E, 0x88, 0x1E, 0x71, 0x53, 0x22, 0xBE, 0x20, 
    0xD3, 0x41, 0xDB, 0x49, 0x95, 0xE4, 0xE4, 0xC7, 0xD2, 0xAE, 
    0xC3, 0x7F, 0x9E, 0xBB, 0x4E, 0xD6, 0x69, 0x9C, 0x93, 0x32, 
    0x4C, 0x46, 0x45, 0x9F, 0xC0, 0x48, 0xBB, 0x96, 0x37, 0x1E, 
    0x3A, 0x89, 0x98, 0xE2, 0xB8, 0x0C, 0x75, 0x5D, 0x5B, 0x01, 
    0xD4, 0x54, 0x0E, 0x58, 0x93, 0xE5, 0x21, 0x47, 0xA3, 0xD5, 
    0x70, 0x1C, 0xF9, 0x83, 0x86, 0x65, 0x9A, 0x4C, 0x9F, 0x55, 
    0x78, 0x12, 0xA1, 0xFF, 0xEF, 0xBA, 0x7F, 0xFB, 0xD4, 0xE8, 
    0xAF, 0x08, 0x00, 0x64, 0x84, 0x17, 0x13, 0x64, 0x3C, 0x91, 
    0x01, 0x82, 0xCC, 0x33, 0x4C, 0x9B, 0x67, 0x2F, 0xBD, 0x94, 
    0x95, 0xB5, 0x2E, 0xFA, 0xA8, 0xD2, 0x62, 0x77, 0xB1, 0xEF, 
    0x8C, 0x8D, 0xB8, 0xA9, 0x1D, 0xFD, 0x03, 0xAC, 0x2D, 0xC1, 
    0x6A, 0x1B, 0x60, 0x39, 0x81, 0xEA, 0x1B, 0x66, 0x0E, 0xA4, 
    0xB2, 0x22, 0x54, 0x8B, 0xCA, 0x47, 0x6F, 0x42, 0xB8, 0xDB, 
    0x49, 0x87, 0xB0, 0x21, 0xCC, 0x13, 0xF1, 0x98, 0x23, 0xCB, 
    0x0D, 0x74, 0x6C, 0x8B, 0xF4, 0xA5, 0xF3, 0x47, 0x2B, 0x62, 
    0x25, 0xB4, 0xE6, 0x09, 0xE5, 0x02, 0x03, 0x01, 0x00, 0x01
};

#endif