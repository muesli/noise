#include "../target/libnoise.h"
#include <string.h>
#include <stdio.h>

int main() {
    Noise_SetLogToStderr(1);

    GoString providerName;
    providerName.p = "ed25519";
    providerName.n = strlen(providerName.p);

    void *kp = Noise_RandomKeyPair(providerName);
    printf("%p\n", kp);

    Noise_Release(kp);

    return 0;
}
