import QtQuick 2.0

Item {
    width: 320; height: 200

    ListView {
        width: 120;
        model: imageModel.len
        delegate: Text {
            text: imageModel.data(index).ID
        }
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.horizontalCenter: parent.horizontalCenter
    }
}
