FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TINI_VERSION
ENV USER_ID=1001
# must match Makefile variable AZURERM_REFSPEC
ENV TERRAFORM_PROVIDER_VERSION=$TERRAFORM_PROVIDER_VERSION
ENV TERRAFORM_PROVIDER_SOURCE=$TERRAFORM_PROVIDER_SOURCE

# Setup Terraform environment
ENV TERRAFORM_VERSION=$TERRAFORM_VERSION
ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/hashicorp/azurerm/${TERRAFORM_PROVIDER_VERSION}/linux_${ARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0
ENV TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/azurerm":{"Protocol":"grpc","ProtocolVersion":5,"Pid":1,"Test":true,"Addr":{"Network":"unix","String":"/tmp/provider"}}}'

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip /tmp
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}
ADD terraform-provider-azurerm-${ARCH} /usr/local/bin/terraform-provider-azurerm
ADD entrypoint.sh /usr/local/bin

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

ADD provider /usr/local/bin/crossplane-jet-azure-provider

USER ${USER_ID}
EXPOSE 8080
ENTRYPOINT ["entrypoint.sh"]
