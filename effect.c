#include "effect.h"

Image *
convolveImage(Image *image, void *data, ExceptionInfo *ex) {
    ConvolveData *d = data;
    return ConvolveImage(image, d->order, d->kernel, ex);
}

Image *
gaussianBlurImage(Image *image, void *data, ExceptionInfo *ex) {
    GaussianBlurData *d = data;
    return GaussianBlurImage(image, d->radius, d->sigma, ex);
}