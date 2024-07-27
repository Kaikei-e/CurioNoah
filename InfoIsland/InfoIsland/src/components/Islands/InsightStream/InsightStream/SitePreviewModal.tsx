import {
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
} from "@chakra-ui/react";

type Props = {
  isOpen: boolean;
  onClose: () => void;
  url: string;
};

export const SitePreviewModal = (props: Props) => {
  const { isOpen, onClose, url } = props;

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="full">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Site Preview</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <iframe
            src={url}
            title="Site Preview"
            width="100%"
            height="100%"
            referrerPolicy="same-origin"
          />
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
