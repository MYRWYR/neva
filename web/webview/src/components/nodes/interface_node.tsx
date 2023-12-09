import { useMemo } from "react";
import classnames from "classnames";
import { Handle, NodeProps, HandleType, Position, useStore } from "reactflow";
import * as src from "../../generated/sourcecode";

export interface IInterfaceNodeProps {
  title: string;
  interface: src.Interface;
}

export function InterfaceNode(props: NodeProps<IInterfaceNodeProps>) {
  const { inports, outports } = useMemo(() => {
    const result = { inports: [], outports: [] };
    if (!props.data.interface.io) {
      return result;
    }
    return {
      inports: Object.entries(props.data.interface.io.in || {}),
      outports: Object.entries(props.data.interface.io.out || {}),
    };
  }, [props.data.interface.io]);

  const isZoomMiddle = useStore((s) => s.transform[2] >= 0.6);
  const isZoomClose = useStore((s) => s.transform[2] >= 1);

  return (
    <div className={"react-flow__node-default"}>
      <Ports
        ports={inports}
        position={Position.Top}
        type="target"
        isVisible={isZoomMiddle}
        areTypesVisible={isZoomClose}
      />
      <div className="nodeBody">
        <div className="nodeName">{props.data.title}</div>
      </div>
      <Ports
        ports={outports}
        position={Position.Bottom}
        type="source"
        isVisible={isZoomMiddle}
        areTypesVisible={isZoomClose}
      />
    </div>
  );
}

function Ports(props: {
  ports: [string, src.Port][];
  position: Position;
  type: HandleType;
  isVisible: boolean;
  areTypesVisible: boolean;
}) {
  if (!props.ports) {
    return null;
  }

  return (
    <div className={classnames("ports", "in", { hidden: !props.isVisible })}>
      {props.ports.map(([portName, portType]) => (
        <Handle
          id={portName}
          type={props.type}
          position={props.position}
          isConnectable={false}
          key={portName}
        >
          {portName}
          {props.areTypesVisible &&
            portType.typeExpr &&
            portType.typeExpr.meta && (
              <span className="portType">
                {" "}
                {(portType.typeExpr.meta as src.Meta).text}
              </span>
            )}
        </Handle>
      ))}
    </div>
  );
}
